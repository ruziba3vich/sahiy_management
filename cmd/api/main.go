package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/sahiy_management/internal/application"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/router"
	"github.com/ruziba3vich/sahiy_management/pkg/config"
	"github.com/ruziba3vich/sahiy_management/pkg/database"
)

// @title           Sahiy Management API
// @version         1.0
// @description     API for managing departments in Sahiy Management system
// @host            localhost:8080
// @BasePath        /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	ctx := context.Background()
	cfg := config.LoadConfig()

	db, err := database.NewPostgresConnection(ctx, &cfg.DBConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := infrastructure.New(db)
	svc := application.New(repo, cfg)

	r := gin.Default()
	router.SetupRoutes(r, svc)

	log.Printf("Server starting on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
