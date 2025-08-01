package controllers

import (
	"net/http"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service ports.UserRepository
}

func NewUserController(service ports.UserRepository) *UserController {
	return &UserController{service: service}
}

// POST /users
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user *domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := uc.service.CreateUser(ctx, user); 
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// PUT /users/:id
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user *domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.service.UpdateUser(ctx, id, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// GET /users/:id
func (uc *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.service.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// GET /users
func (uc *UserController) ListUsers(ctx *gin.Context) {
	users, err := uc.service.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// DELETE /users/:id
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := uc.service.DeleteUser(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
