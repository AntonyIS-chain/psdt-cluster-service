package controllers


import (
	"net/http"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service ports.UserService
}

func NewUserController(service ports.UserService) *UserController {
	return &UserController{service: service}
}

// POST /auth/login
func (uc *UserController) Authenticate(ctx *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&creds); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.service.Authenticate(ctx,creds.Username, creds.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// GET /users/:username
func (uc *UserController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := uc.service.GetUserByUsername(ctx,username)
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

// POST /users/invite
func (uc *UserController) InviteUser(ctx *gin.Context) {
	var req struct {
		Email string `json:"email"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := uc.service.InviteUser(ctx,req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"invite_token": token})
}

// POST /users/register
func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var signup domain.UserSignup
	if err := ctx.ShouldBindJSON(&signup); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.service.RegisterUser(ctx,signup)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
