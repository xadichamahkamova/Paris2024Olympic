package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"user-service/logger"
	"user-service/token"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/userpb"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db  *sql.DB
	rds *redis.Client
}

func NewPostgresUserRepo(db *sql.DB, rds *redis.Client) UserRepository {
	return &UserRepo{
		db:  db,
		rds: rds,
	}
}

func (u *UserRepo) Register(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &pb.User{}
	now := time.Now().Format(time.RFC3339)

	// Parolni hashlash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to hash password", logrus.Fields{
			"username": req.Username,
			"error":    err,
		})
		return &pb.CreateUserResponse{Success: false, Message: "Failed to hash password"}, err
	}

	query := 
	`INSERT INTO users (username, password, role, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING id, username, role, created_at, updated_at`
	err = u.db.QueryRow(query, 
		req.Username, 
		string(hashedPassword), 
		req.Role, 
		now, now,
	).Scan(&user.Id, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		logger.Error("Failed to create user", logrus.Fields{
			"username": req.Username,
			"error":    err,
		})
		return &pb.CreateUserResponse{Success: false, Message: "Failed to create user"}, err
	}

	logger.Info("User created successfully", logrus.Fields{
		"username": req.Username,
		"user_id": user.Id,
	})

	return &pb.CreateUserResponse{
		Success: true,
		Message: "User created successfully",
		User:    user,
	}, nil
}

func (s *UserRepo) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	userResp, err := s.GetUserByFilter(ctx, &pb.UserFilter{
		Username: req.Username,
	})

	if err != nil || len(userResp.Users) == 0 {
		logger.Warn("Invalid login attempt", logrus.Fields{
			"username": req.Username,
		})
		return &pb.LoginResponse{
			Success: false,
			Message: "Invalid credentials",
		}, nil
	}

	user := userResp.Users[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		logger.Warn("Invalid password attempt", logrus.Fields{
			"username": req.Username,
		})
		return &pb.LoginResponse{
			Success: false,
			Message: "Invalid credentials",
		}, nil
	}

	accessToken, refreshToken, err := token.CreateTokens(user)
	if err != nil {
		logger.Error("Failed to create tokens", logrus.Fields{
			"username": req.Username,
			"error":    err,
		})
		return nil, err
	}

	err = s.rds.Set(ctx, refreshToken, user.Id, time.Hour*24).Err()
	if err != nil {
		logger.Error("Failed to set refresh token in Redis", logrus.Fields{
			"refresh_token": refreshToken,
			"error":        err,
		})
		return nil, err
	}

	logger.Info("Login successful", logrus.Fields{
		"username": req.Username,
	})

	return &pb.LoginResponse{
		Success:      true,
		Message:      "Login successful",
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserRepo) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	userId, err := s.rds.Get(ctx, req.RefreshToken).Result()
	if err != nil {
		logger.Warn("Invalid refresh token", logrus.Fields{
			"refresh_token": req.RefreshToken,
			"error":        err,
		})
		return &pb.RefreshTokenResponse{
			Success: false,
			Message: "Invalid refresh token",
		}, nil
	}

	userResp, err := s.GetUserById(ctx, &pb.GetUserRequest{Id: userId})
	if err != nil {
		logger.Error("Failed to get user by ID", logrus.Fields{
			"user_id": userId,
			"error":   err,
		})
		return nil, err
	}
	user := userResp.User

	accessToken, refreshToken, err := token.CreateTokens(user)
	if err != nil {
		logger.Error("Failed to create tokens", logrus.Fields{
			"user_id": user.Id,
			"error":   err,
		})
		return nil, err
	}

	err = s.rds.Set(ctx, refreshToken, user.Id, time.Hour*24).Err()
	if err != nil {
		logger.Error("Failed to set new refresh token in Redis", logrus.Fields{
			"refresh_token": refreshToken,
			"error":        err,
		})
		return nil, err
	}

	logger.Info("Token refreshed successfully", logrus.Fields{
		"username": user.Username,
	})

	return &pb.RefreshTokenResponse{
		Success:      true,
		Message:      "Token refreshed successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	now := time.Now().Format(time.RFC3339)

	var err error
	if req.User.Password != "" {
		// Agar parol yangilanayotgan bo'lsa, uni hashlash
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
		if err != nil {
			logger.Error("Failed to hash password for update", logrus.Fields{
				"user_id": req.User.Id,
				"error":   err,
			})
			return &pb.UpdateUserResponse{Success: false, Message: "Failed to hash password"}, err
		}
		_, err = u.db.Exec(
			"UPDATE users SET username = $1, role = $2, password = $3, updated_at = $4 WHERE id = $5",
			req.User.Username, req.User.Role, string(hashedPassword), now, req.User.Id,
		)
		if err != nil {
			logger.Error("Failed to update user", logrus.Fields{
				"user_id": req.User.Id,
				"error":   err,
			})
			return nil, err
		}
	} else {
		_, err = u.db.Exec(
			"UPDATE users SET username = $1, role = $2, updated_at = $3 WHERE id = $4",
			req.User.Username, req.User.Role, now, req.User.Id,
		)
	}

	if err != nil {
		logger.Error("Failed to update user", logrus.Fields{
			"user_id": req.User.Id,
			"error":   err,
		})
		return &pb.UpdateUserResponse{Success: false, Message: "Failed to update user"}, err
	}

	req.User.UpdatedAt = now

	logger.Info("User updated successfully", logrus.Fields{
		"user_id": req.User.Id,
		"username": req.User.Username,
	})

	return &pb.UpdateUserResponse{
		Success: true,
		Message: "User updated successfully",
		User:    req.User,
	}, nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, err := u.db.Exec(
		"UPDATE users SET deleted_at = $1 WHERE id = $2",
		time.Now().Unix(), req.Id,
	)

	if err != nil {
		logger.Error("Failed to delete user", logrus.Fields{
			"user_id": req.Id,
			"error":   err,
		})
		return &pb.DeleteUserResponse{Success: false, Message: "Failed to delete user"}, err
	}

	logger.Info("User deleted successfully", logrus.Fields{
		"user_id": req.Id,
	})

	return &pb.DeleteUserResponse{
		Success: true,
		Message: "User deleted successfully",
	}, nil
}


func (u *UserRepo) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user := &pb.User{}
	err := u.db.QueryRow(
		"SELECT id, username, role, created_at, updated_at FROM users WHERE id = $1 AND deleted_at = 0",
		req.Id,
	).Scan(&user.Id, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("User not found", logrus.Fields{
				"user_id": req.Id,
			})
			return &pb.GetUserResponse{Success: false, Message: "User not found"}, nil
		}
		logger.Error("Failed to retrieve user", logrus.Fields{
			"user_id": req.Id,
			"error":   err,
		})
		return &pb.GetUserResponse{Success: false, Message: "Failed to get user"}, err
	}

	logger.Info("User retrieved successfully", logrus.Fields{
		"user_id": req.Id,
		"username": user.Username,
	})

	return &pb.GetUserResponse{
		Success: true,
		Message: "User retrieved successfully",
		User:    user,
	}, nil
}

func (u *UserRepo) GetUsers(ctx context.Context, req *pb.Void) (*pb.GetUsersResponse, error) {
	rows, err := u.db.Query("SELECT id, username, role, created_at, updated_at FROM users WHERE deleted_at = 0")
	if err != nil {
		logger.Error("Failed to retrieve users", logrus.Fields{
			"error": err,
		})
		return &pb.GetUsersResponse{Success: false, Message: "Failed to get users"}, err
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		user := &pb.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			logger.Error("Failed to scan user row", logrus.Fields{
				"error": err,
			})
			return &pb.GetUsersResponse{Success: false, Message: "Failed to scan user"}, err
		}
		users = append(users, user)
	}

	logger.Info("Users retrieved successfully", logrus.Fields{
		"count": len(users),
	})

	return &pb.GetUsersResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Users:   users,
	}, nil
}

func (u *UserRepo) GetUserByFilter(ctx context.Context, req *pb.UserFilter) (*pb.GetUsersResponse, error) {
	query := "SELECT id, username, password, role, created_at, updated_at FROM users WHERE deleted_at = 0"
	args := []interface{}{}

	if req.Username != "" {
		query += " AND username LIKE $1"
		args = append(args, "%"+req.Username+"%")
	}

	if req.Role != "" {
		query += " AND role = $" + fmt.Sprint(len(args)+1)
		args = append(args, req.Role)
	}

	rows, err := u.db.Query(query, args...)
	if err != nil {
		logger.Error("Failed to retrieve users by filter", logrus.Fields{
			"error": err,
		})
		return &pb.GetUsersResponse{Success: false, Message: "Failed to get users"}, err
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		user := &pb.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			logger.Error("Failed to scan user row", logrus.Fields{
				"error": err,
			})
			return &pb.GetUsersResponse{Success: false, Message: "Failed to scan user"}, err
		}
		users = append(users, user)
	}

	logger.Info("Users retrieved successfully by filter", logrus.Fields{
		"username": req.Username,
		"role":     req.Role,
		"count":    len(users),
	})

	return &pb.GetUsersResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Users:   users,
	}, nil
}
