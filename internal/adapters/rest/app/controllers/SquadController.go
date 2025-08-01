package controllers

import (
	"net/http"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type SquadController struct {
	service ports.SquadRepository
}

func NewSquadController(service ports.SquadRepository) *SquadController {
	return &SquadController{service: service}
}

// POST /squads
func (sc *SquadController) CreateSquad(ctx *gin.Context) {
	var squad *domain.Squad
	if err := ctx.ShouldBindJSON(&squad); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := sc.service.CreateSquad(ctx, squad)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// PUT /squads/:id
func (sc *SquadController) UpdateSquad(ctx *gin.Context) {
	id := ctx.Param("id")
	var squad *domain.Squad
	if err := ctx.ShouldBindJSON(&squad); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := sc.service.UpdateSquad(ctx, id, squad); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// GET /squads/:id
func (sc *SquadController) GetSquadByID(ctx *gin.Context) {
	id := ctx.Param("id")
	squad, err := sc.service.GetSquadByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, squad)
}

// GET /squads
func (sc *SquadController) ListSquads(ctx *gin.Context) {
	squads, err := sc.service.ListSquads(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, squads)
}

// DELETE /squads/:id
func (sc *SquadController) DeleteSquad(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := sc.service.DeleteSquad(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// POST /squads/:squadID/users/:userID
func (sc *SquadController) AddUserToSquad(ctx *gin.Context) {
	squadID := ctx.Param("squadID")
	userID := ctx.Param("userID")

	if err := sc.service.AddUserToSquad(ctx, squadID, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// DELETE /squads/:squadID/users/:userID
func (sc *SquadController) RemoveUserFromSquad(ctx *gin.Context) {
	squadID := ctx.Param("squadID")
	userID := ctx.Param("userID")

	if err := sc.service.RemoveUserFromSquad(ctx, squadID, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// GET /squads/:squadID/users
func (sc *SquadController) ListUsersInSquad(ctx *gin.Context) {
	squadID := ctx.Param("squadID")

	users, err := sc.service.ListUsersInSquad(ctx, squadID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
