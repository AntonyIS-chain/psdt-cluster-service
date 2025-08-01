package controllers

import (
	"net/http"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type TribeController struct {
	service ports.TribeRepository
}

func NewTribeController(service ports.TribeRepository) *TribeController {
	return &TribeController{service: service}
}

// GET /tribes/:id
func (tc *TribeController) GetTribeByID(ctx *gin.Context) {
	id := ctx.Param("id")
	tribe, err := tc.service.GetTribeByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tribe)
}

// GET /tribes
func (tc *TribeController) ListTribes(ctx *gin.Context) {
	tribes, err := tc.service.ListTribes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tribes)
}

// POST /tribes
func (tc *TribeController) CreateTribe(ctx *gin.Context) {
	var tribe *domain.Tribe
	if err := ctx.ShouldBindJSON(&tribe); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := tc.service.CreateTribe(ctx, tribe)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// PUT /tribes/:id
func (tc *TribeController) UpdateTribe(ctx *gin.Context) {
	id := ctx.Param("id")
	var updates *domain.Tribe
	if err := ctx.ShouldBindJSON(&updates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := tc.service.UpdateTribe(ctx, id, updates); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// DELETE /tribes/:id
func (tc *TribeController) DeleteTribe(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := tc.service.DeleteTribe(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// POST /tribes/:tribeID/squads/:squadID
func (tc *TribeController) AddSquadToTribe(ctx *gin.Context) {
	tribeID := ctx.Param("tribeID")
	squadID := ctx.Param("squadID")

	err := tc.service.AddSquadToTribe(ctx, tribeID, squadID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// DELETE /tribes/:tribeID/squads/:squadID
func (tc *TribeController) RemoveSquadFromTribe(ctx *gin.Context) {
	tribeID := ctx.Param("tribeID")
	squadID := ctx.Param("squadID")

	err := tc.service.RemoveSquadFromTribe(ctx, tribeID, squadID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// GET /tribes/:tribeID/squads
func (tc *TribeController) ListSquadsInTribe(ctx *gin.Context) {
	tribeID := ctx.Param("tribeID")

	squads, err := tc.service.ListSquadsInTribe(ctx, tribeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, squads)
}
