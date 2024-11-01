package repository

import (
	"context"
	"testing"
	"time"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/userpb"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func setupTest(t *testing.T) (*UserRepo, sqlmock.Sqlmock, *redis.Client, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a mock redis server", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	repo := NewPostgresUserRepo(db, rdb).(*UserRepo)

	return repo, mock, rdb, func() {
		db.Close()
		mr.Close()
	}
}

func TestRegister(t *testing.T) {
	repo, mock, _, teardown := setupTest(t)
	defer teardown()

	ctx := context.Background()
	req := &pb.CreateUserRequest{
		Username: "mongosh",
		Password: "1001",
		Role:     "admin",
	}

	mock.ExpectQuery("INSERT INTO users").
		WithArgs(req.Username, sqlmock.AnyArg(), req.Role, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "role", "created_at", "updated_at"}).
			AddRow("1", req.Username, req.Role, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339)))

	resp, err := repo.Register(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "User created successfully", resp.Message)
	assert.Equal(t, req.Username, resp.User.Username)
	assert.Equal(t, req.Role, resp.User.Role)
}

func TestLogin(t *testing.T) {
	repo, mock, rdb, teardown := setupTest(t)
	defer teardown()

	ctx := context.Background()
	req := &pb.LoginRequest{
		Username: "mongosh",
		Password: "1001",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs("%" + req.Username + "%").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "role", "created_at", "updated_at"}).
			AddRow("1", req.Username, string(hashedPassword), "user", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339)))

	resp, err := repo.Login(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "Login successful", resp.Message)
	assert.NotEmpty(t, resp.AccessToken)
	assert.NotEmpty(t, resp.RefreshToken)

	// Check if refresh token is set in Redis
	val, err := rdb.Get(ctx, resp.RefreshToken).Result()
	assert.NoError(t, err)
	assert.Equal(t, "1", val)
}

func TestRefreshToken(t *testing.T) {
	repo, mock, rdb, teardown := setupTest(t)
	defer teardown()

	ctx := context.Background()
	refreshToken := "testRefreshToken"
	userId := "1"

	// Set refresh token in Redis
	err := rdb.Set(ctx, refreshToken, userId, time.Hour).Err()
	assert.NoError(t, err)

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "role", "created_at", "updated_at"}).
			AddRow(userId, "testuser", "user", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339)))

	resp, err := repo.RefreshToken(ctx, &pb.RefreshTokenRequest{RefreshToken: refreshToken})

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "Token refreshed successfully", resp.Message)
	assert.NotEmpty(t, resp.AccessToken)
	assert.NotEmpty(t, resp.RefreshToken)

	// Check if new refresh token is set in Redis
	val, err := rdb.Get(ctx, resp.RefreshToken).Result()
	assert.NoError(t, err)
	assert.Equal(t, userId, val)
}

func TestUpdateUser(t *testing.T) {
	repo, mock, _, teardown := setupTest(t)
	defer teardown()

	ctx := context.Background()
	req := &pb.UpdateUserRequest{
		User: &pb.User{
			Id:       "1",
			Username: "mongosh)",
			Role:     "admin",
			Password: "1001",
		},
	}

	mock.ExpectExec("UPDATE users SET").
		WithArgs(req.User.Username, req.User.Role, sqlmock.AnyArg(), sqlmock.AnyArg(), req.User.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.UpdateUser(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "User updated successfully", resp.Message)
	assert.Equal(t, req.User.Username, resp.User.Username)
	assert.Equal(t, req.User.Role, resp.User.Role)
}

func TestDeleteUser(t *testing.T) {
	repo, mock, _, teardown := setupTest(t)
	defer teardown()

	ctx := context.Background()
	req := &pb.DeleteUserRequest{
		Id: "1",
	}

	mock.ExpectExec("UPDATE users SET deleted_at").
		WithArgs(sqlmock.AnyArg(), req.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.DeleteUser(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "User deleted successfully", resp.Message)
}

func TestGetUserById(t *testing.T) {
	repo, mock, _, teardown := setupTest(t)
	defer teardown()

	ctx := context.Background()
	req := &pb.GetUserRequest{
		Id: "1",
	}

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "role", "created_at", "updated_at"}).
			AddRow(req.Id, "testuser", "user", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339)))

	resp, err := repo.GetUserById(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "User retrieved successfully", resp.Message)
	assert.Equal(t, req.Id, resp.User.Id)
	assert.Equal(t, "testuser", resp.User.Username)
}

func TestGetUsers(t *testing.T) {
	repo, mock, _, teardown := setupTest(t)
	defer teardown()

	ctx := context.Background()

	mock.ExpectQuery("SELECT (.+) FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "role", "created_at", "updated_at"}).
			AddRow("1", "user1", "user", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339)).
			AddRow("2", "user2", "admin", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339)))

	resp, err := repo.GetUsers(ctx, &pb.Void{})

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "Users retrieved successfully", resp.Message)
	assert.Len(t, resp.Users, 2)
}

func TestGetUserByFilter(t *testing.T) {
	repo, mock, _, teardown := setupTest(t)
	defer teardown()

	ctx := context.Background()
	req := &pb.UserFilter{
		Username: "mongosh",
		Role:     "admin",
	}

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs("%"+req.Username+"%", req.Role).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "role", "created_at", "updated_at"}).
			AddRow("1", "testuser", "hashedpassword", "user", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339)))

	resp, err := repo.GetUserByFilter(ctx, req)

	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "Users retrieved successfully", resp.Message)
	assert.Len(t, resp.Users, 1)
	assert.Equal(t, "testuser", resp.Users[0].Username)
	assert.Equal(t, "user", resp.Users[0].Role)
}
