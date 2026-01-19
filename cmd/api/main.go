package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	appDept "github.com/ruziba3vich/sahiy_management/internal/application/department"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/handler"
	"github.com/ruziba3vich/sahiy_management/pkg/database"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ruziba3vich/sahiy_management/docs"
)

// @title           Sahiy Management API
// @version         1.0
// @description     API for managing departments in Sahiy Management system
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     dbPort,
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "sahiy_management"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	db, err := database.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	deptRepo := postgres.NewDepartmentRepository(db)
	deptService := appDept.NewService(deptRepo)
	deptHandler := handler.NewDepartmentHandler(deptService)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	deptHandler.RegisterRoutes(api)

	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
