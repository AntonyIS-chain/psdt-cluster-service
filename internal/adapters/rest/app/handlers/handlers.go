package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AntonyIS-chain/psdt-cluster-service/config"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/adapters/rest/app/controllers"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGinRoutes(
	clusterSvc ports.ClusterRepository,
	tribeSvc ports.TribeRepository,
	squadSvc ports.SquadRepository,
	userSvc ports.UserRepository,
	config *config.Config,
) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize only relevant controllers
	clusterController := controllers.NewClusterController(clusterSvc)
	tribeController := controllers.NewTribeController(tribeSvc)
	squadController := controllers.NewSquadController(squadSvc)
	userController := controllers.NewUserController(userSvc)

	// --- User Routes ---

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to PSDT Cluster Service",
		})
	})

	router.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Health Check",
		})
	})

	// --- Cluster Routes ---
	clusterGroup := router.Group("/clusters")
	{
		clusterGroup.GET("/", clusterController.ListClusters)
		clusterGroup.GET("/:id", clusterController.GetClusterByID)
		clusterGroup.POST("/", clusterController.CreateCluster)
		clusterGroup.PUT("/:id", clusterController.UpdateCluster)
		clusterGroup.DELETE("/:id", clusterController.DeleteCluster)

		tribeGroup := clusterGroup.Group("/tribes/:clusterID/tribes")

		{
			tribeGroup.POST("/:clusterID/tribes/:tribeID", clusterController.AddTribeToCluster)
			tribeGroup.DELETE("/:clusterID/tribes/:tribeID", clusterController.RemoveTribeFromCluster)
			tribeGroup.GET("/:clusterID/tribes", clusterController.ListTribesInCluster)
		}
		
	}

	// --- Tribe Routes ---
	tribeRoutes := router.Group("/v1/api/tribes")
	{
		tribeRoutes.POST("/", tribeController.CreateTribe)
		tribeRoutes.GET("/", tribeController.ListTribes)
		tribeRoutes.GET("/:id", tribeController.GetTribeByID)
		tribeRoutes.PUT("/:id", tribeController.UpdateTribe)
		tribeRoutes.DELETE("/:id", tribeController.DeleteTribe)
	}

	// --- Squad Routes ---
	squadGroup := router.Group("/v1/api/squads")
	{
		squadGroup.POST("/", squadController.CreateSquad)
		squadGroup.GET("/", squadController.ListSquads)
		squadGroup.GET("/:id", squadController.GetSquadByID)
		squadGroup.PUT("/:id", squadController.UpdateSquad)
		squadGroup.DELETE("/:id", squadController.DeleteSquad)

		usersGroup := clusterGroup.Group("/users/:clusterID/tribes")

		{
			usersGroup.POST("/:squadID/users/:userID", squadController.AddUserToSquad)
			usersGroup.DELETE("/:squadID/users/:userID", squadController.RemoveUserFromSquad)
			usersGroup.GET("/:squadID/users", squadController.ListUsersInSquad)
		}
		
	}

	userRoutes := router.Group("/v1/api/users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/", userController.ListUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}

	// --- Start Server ---
	log.Println("Starting server on port", config.SERVER_PORT)
	router.Run(fmt.Sprintf(":%s", config.SERVER_PORT))
}
