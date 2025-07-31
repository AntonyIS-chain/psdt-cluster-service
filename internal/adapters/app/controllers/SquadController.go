package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
)

type SquadController struct {
	service ports.SquadService
}

func NewSquadController(service ports.SquadService) *SquadController {
	return &SquadController{service: service}
}

// GET /squads/:id
func (sc *SquadController) GetSquadByID(ctx *gin.Context) {
	id := ctx.Param("id")
	squad, err := sc.service.GetSquadByID(ctx,id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, squad)
}

// GET /tribes/:tribeID/squads
func (sc *SquadController) ListSquads(ctx *gin.Context) {
	tribeID := ctx.Param("tribeID")
	squads, err := sc.service.ListSquads(ctx,tribeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, squads)
}

// POST /squads
func (sc *SquadController) CreateSquad(ctx *gin.Context) {
	var squad domain.Squad
	if err := ctx.ShouldBindJSON(&squad); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := sc.service.CreateSquad(ctx,squad)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// PUT /squads/:id
func (sc *SquadController) UpdateSquad(ctx *gin.Context) {
	id := ctx.Param("id")
	var updates domain.Squad
	if err := ctx.ShouldBindJSON(&updates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sc.service.UpdateSquad(ctx,id, updates); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// DELETE /squads/:id
func (sc *SquadController) DeleteSquad(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := sc.service.DeleteSquad(ctx,id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// POST /squads/:squadID/users
func (sc *SquadController) AddUserToSquad(ctx *gin.Context) {
	var user domain.SquadUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := sc.service.AddUserToSquad(ctx,user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// DELETE /squads/:squadID/users/:userID
func (sc *SquadController) RemoveUserFromSquad(ctx *gin.Context) {
	squadID := ctx.Param("squadID")
	userID := ctx.Param("userID")

	if err := sc.service.RemoveUserFromSquad(ctx,userID, squadID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// GET /squads/:squadID/users
func (sc *SquadController) ListUsersInSquad(ctx *gin.Context) {
	squadID := ctx.Param("squadID")
	users, err := sc.service.ListUsersInSquad(ctx,squadID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// GET /users/:userID/squads
func (sc *SquadController) GetSquadsForUser(ctx *gin.Context) {
	userID := ctx.Param("userID")
	squads, err := sc.service.GetSquadsForUser(ctx,userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, squads)
}
