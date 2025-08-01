package controllers

import (
	"net/http"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type ClusterController struct {
	service ports.ClusterRepository
}

func NewClusterController(service ports.ClusterRepository) *ClusterController {
	return &ClusterController{service: service}
}

// GET /clusters/:id
func (cc *ClusterController) GetClusterByID(ctx *gin.Context) {
	id := ctx.Param("id")
	cluster, err := cc.service.GetClusterByID(ctx, id)
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
	var cluster *domain.Cluster
	if err := ctx.ShouldBindJSON(&cluster); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := cc.service.CreateCluster(ctx, cluster)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": id})
}

// PUT /clusters/:id
func (cc *ClusterController) UpdateCluster(ctx *gin.Context) {
	id := ctx.Param("id")
	var updates *domain.Cluster
	if err := ctx.ShouldBindJSON(&updates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cc.service.UpdateCluster(ctx, id, updates); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// DELETE /clusters/:id
func (cc *ClusterController) DeleteCluster(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := cc.service.DeleteCluster(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// POST /clusters/:clusterID/tribes/:tribeID
func (cc *ClusterController) AddTribeToCluster(ctx *gin.Context) {
	clusterID := ctx.Param("clusterID")
	tribeID := ctx.Param("tribeID")

	err := cc.service.AddTribeToCluster(ctx, clusterID, tribeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// DELETE /clusters/:clusterID/tribes/:tribeID
func (cc *ClusterController) RemoveTribeFromCluster(ctx *gin.Context) {
	clusterID := ctx.Param("clusterID")
	tribeID := ctx.Param("tribeID")

	err := cc.service.RemoveTribeFromCluster(ctx, clusterID, tribeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// GET /clusters/:clusterID/tribes
func (cc *ClusterController) ListTribesInCluster(ctx *gin.Context) {
	clusterID := ctx.Param("clusterID")

	tribes, err := cc.service.ListTribesInCluster(ctx, clusterID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tribes)
}
