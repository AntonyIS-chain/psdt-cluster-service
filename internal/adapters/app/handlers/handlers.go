package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AntonyIS-chain/psdt-core-service/config"
	"github.com/AntonyIS-chain/psdt-core-service/internal/adapters/app/controllers"
	"github.com/AntonyIS-chain/psdt-core-service/internal/core/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGinRoutes(
	clusterSvc ports.ClusterService,
	tribeSvc ports.TribeService,
	squadSvc ports.SquadService,
	userSvc ports.UserService,
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
	userController := controllers.NewUserController(userSvc)
	clusterController := controllers.NewClusterController(clusterSvc)
	tribeController := controllers.NewTribeController(tribeSvc)
	squadController := controllers.NewSquadController(squadSvc)

	// --- User Routes ---
	userRoutes := router.Group("/v1/api/users")
	{
		userRoutes.GET("/:id", userController.GetUserByUsername)
		userRoutes.GET("/", userController.ListUsers)
		userRoutes.POST("/", userController.RegisterUser)
		userRoutes.POST("/invite", userController.InviteUser)
		userRoutes.POST("/authenticate", userController.Authenticate)
		userRoutes.GET("/username/:username", userController.GetUserByUsername)
	}

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
	clusterRoutes := router.Group("/v1/api/clusters")
	{
		clusterRoutes.POST("/", clusterController.CreateCluster)
		clusterRoutes.GET("/", clusterController.ListClusters)
		clusterRoutes.GET("/:id", clusterController.GetClusterByID)
		clusterRoutes.PUT("/:id", clusterController.UpdateCluster)
		clusterRoutes.DELETE("/:id", clusterController.DeleteCluster)
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
	squadRoutes := router.Group("/v1/api/squads")
	{
		squadRoutes.POST("/", squadController.CreateSquad)
		squadRoutes.GET("/:id", squadController.GetSquadByID)
		squadRoutes.PUT("/:id", squadController.UpdateSquad)
		squadRoutes.DELETE("/:id", squadController.DeleteSquad)

		squadRoutes.POST("/add-user", squadController.AddUserToSquad)
		squadRoutes.DELETE("/remove-user/:user_id/:squad_id", squadController.RemoveUserFromSquad)
		squadRoutes.GET("/users/:squad_id", squadController.ListUsersInSquad)
		squadRoutes.GET("/by-user/:user_id", squadController.GetSquadsForUser)
	}

	// --- Start Server ---
	log.Println("Starting server on port", config.SERVER_PORT)
	router.Run(fmt.Sprintf(":%s", config.SERVER_PORT))
}
