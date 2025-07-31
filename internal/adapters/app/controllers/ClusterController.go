package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
)

type ClusterController struct {
	service ports.ClusterService
}

func NewClusterController(service ports.ClusterService) *ClusterController {
	return &ClusterController{service: service}
}

// GET /clusters/:id
func (cc *ClusterController) GetClusterByID(ctx *gin.Context) {
	id := ctx.Param("id")
	cluster, err := cc.service.GetClusterByID(ctx,id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cluster)
}

// GET /clusters
func (cc *ClusterController) ListClusters(ctx *gin.Context) {
	clusters, err := cc.service.ListClusters(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, clusters)
}

// POST /clusters
func (cc *ClusterController) CreateCluster(ctx *gin.Context) {
	var cluster domain.Cluster
	if err := ctx.ShouldBindJSON(&cluster); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := cc.service.CreateCluster(ctx,cluster)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// PUT /clusters/:id
func (cc *ClusterController) UpdateCluster(ctx *gin.Context) {
	id := ctx.Param("id")
	var updates domain.Cluster
	if err := ctx.ShouldBindJSON(&updates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.service.UpdateCluster(ctx,id, updates); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// DELETE /clusters/:id
func (cc *ClusterController) DeleteCluster(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := cc.service.DeleteCluster(ctx,id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
