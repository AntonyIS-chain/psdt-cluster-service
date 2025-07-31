package controllers

import (
	"net/http"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type TribeController struct {
	tribeService ports.TribeService
}

func NewTribeController(tribeService ports.TribeService) *TribeController {
	return &TribeController{
		tribeService: tribeService,
	}
}

func (tc *TribeController) GetTribeByID(c *gin.Context) {
	id := c.Param("id")
	tribe, err := tc.tribeService.GetTribeByID(c,id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tribe)
}

func (tc *TribeController) ListTribes(c *gin.Context) {
	clusterID := c.Query("clusterId")
	tribes, err := tc.tribeService.ListTribes(c,clusterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tribes)
}

func (tc *TribeController) CreateTribe(c *gin.Context) {
	var tribe domain.Tribe
	if err := c.ShouldBindJSON(&tribe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := tc.tribeService.CreateTribe(c,tribe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (tc *TribeController) UpdateTribe(c *gin.Context) {
	id := c.Param("id")
	var updates domain.Tribe
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.tribeService.UpdateTribe(c,id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (tc *TribeController) DeleteTribe(c *gin.Context) {
	id := c.Param("id")
	if err := tc.tribeService.DeleteTribe(c,id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
