package handler 

import (
	"context"
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/userpb"
	"api-gateway/logger"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Router /auth/register [post]
// @Summary REGISTER USER
// @Description This method registers a new user
// @Security BearerAuth
// @Tags AUTH
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) RegisterUser(c *gin.Context) {

	req := pb.CreateUserRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("RegisterUser: Failed to bind JSON: ", err)
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.Register(context.Background(), &req)
	if err != nil {
		logger.Error("RegisterUser: Failed to register user: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("RegisterUser: User registered successfully: ", logrus.Fields{
		"id":   resp.User.Id,
		"role": resp.User.Role,
	})
	c.JSON(200, resp)
}

// @Router /auth/login [post]
// @Summary LOGIN USER
// @Description This method logs in a user
// @Security BearerAuth
// @Tags AUTH
// @Accept json
// @Produce json
// @Param user body models.LoginRequest true "User"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) LoginUser(c *gin.Context) {

	req := pb.LoginRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("LoginUser: Failed to bind JSON: ", err)
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.Login(context.Background(), &req)
	if err != nil {
		logger.Error("LoginUser: Failed to login user: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("LoginUser: User logged in successfully: ", logrus.Fields{
		"id":    resp.User.Id,
		"token": resp.AccessToken,
	})
	c.JSON(200, resp)
}

// @Router /auth/refresh [post]
// @Summary REFRESH TOKEN
// @Description This method refreshes the authentication token
// @Security BearerAuth
// @Tags AUTH
// @Accept json
// @Produce json
// @Param token body models.RefreshTokenRequest true "Token"
// @Success 200 {object} models.RefreshTokenResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) RefreshToken(c *gin.Context) {

	req := pb.RefreshTokenRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("RefreshToken: Failed to bind JSON: ", err)
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.RefreshToken(context.Background(), &req)
	if err != nil {
		logger.Error("RefreshToken: Failed to refresh token: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("RefreshToken: Token refreshed successfully")
	c.JSON(200, resp)
}

// @Router /users/{id} [put]
// @Summary UPDATE USER
// @Description This method updates a user
// @Security BearerAuth
// @Tags USER
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param user body models.UpdateUserRequest true "User"
// @Success 200 {object} models.User
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) UpdateUser(c *gin.Context) {

	req := pb.UpdateUserRequest{
        User: &pb.User{},
    }
    req.User.Id = c.Param("id")

	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateUser: Failed to bind JSON for user ID ", logrus.Fields{
			"id": req.User.Id,
		})
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.UpdateUser(context.Background(), &req)
	if err != nil {
		logger.Error("UpdateUser: Failed to update user with ID ", logrus.Fields{
			"id": req.User.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("UpdateUser: User updated successfully: ", logrus.Fields{
		"id":   resp.User.Id,
		"name": resp.User.Username,
	})
	c.JSON(200, resp)
}

// @Router /users/{id} [delete]
// @Summary DELETE USER
// @Description This method deletes a user
// @Security BearerAuth
// @Tags USER
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.DeleteUserResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) DeleteUser(c *gin.Context) {

	req := pb.DeleteUserRequest{}
	req.Id = c.Param("id")
	resp, err := h.Service.DeleteUser(context.Background(), &req)
	if err != nil {
		logger.Error("DeleteUser: Failed to delete user with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}

	logger.Info("DeleteUser: User deleted successfully")
	c.JSON(200, resp)
}

// @Router /users/{id} [get]
// @Summary GET USER
// @Description This method gets a user by ID
// @Security BearerAuth
// @Tags USER
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.GetUserResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetUserById(c *gin.Context) {

	req := pb.GetUserRequest{}
	req.Id = c.Param("id")
	resp, err := h.Service.GetUserById(context.Background(), &req)
	if err != nil {
		logger.Error("GetUserById: Failed to get user with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("GetUserById: User retrieved successfully: ", logrus.Fields{
		"id":   resp.User.Id,
		"name": resp.User.Username,
	})
	c.JSON(200, resp)
}

// @Router /users [get]
// @Summary GET USERS
// @Description This method gets all users
// @Security BearerAuth
// @Tags USER
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUsersResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetUsers(c *gin.Context) {

	resp, err := h.Service.GetUsers(context.Background(), &pb.Void{})
	if err != nil {
		logger.Error("GetUsers: Failed to get users: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}

	logger.Info("GetUsers: Users retrieved successfully")
	c.JSON(200, resp)
}

// @Router /users/filter [post]
// @Summary GET USERS BY FILTER
// @Description This method gets users by filter
// @Security BearerAuth
// @Tags USER
// @Accept json
// @Produce json
// @Param filter body models.UserFilter true "Filter"
// @Success 200 {object} models.GetUsersResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetUserByFilter(c *gin.Context) {
	
	req := pb.UserFilter{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("GetUserByFilter: Failed to bind JSON: ", err)
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.GetUserByFilter(context.Background(), &req)
	if err != nil {
		logger.Error("GetUserByFilter: Failed to get users by filter: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("GetUserByFilter: Users retrieved successfully by filter")
	c.JSON(200, resp)
}
