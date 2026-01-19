package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	appAuth "github.com/ruziba3vich/sahiy_management/internal/application/auth"
	appDept "github.com/ruziba3vich/sahiy_management/internal/application/department"
	appPriv "github.com/ruziba3vich/sahiy_management/internal/application/privilege"
	appSec "github.com/ruziba3vich/sahiy_management/internal/application/section"
	appTask "github.com/ruziba3vich/sahiy_management/internal/application/task"
	appTH "github.com/ruziba3vich/sahiy_management/internal/application/taskhistory"
	appTS "github.com/ruziba3vich/sahiy_management/internal/application/taskstatus"
	appUser "github.com/ruziba3vich/sahiy_management/internal/application/user"
	"github.com/ruziba3vich/sahiy_management/internal/infrastructure/postgres"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/handler"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/middleware"
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

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	ctx := context.Background()

	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
	jwtExpiryHours, _ := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))

	dbConfig := database.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     dbPort,
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "sahiy_management"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	jwtSecret := getEnv("JWT_SECRET", "your-256-bit-secret-change-in-production")

	db, err := database.NewPostgresConnection(ctx, &dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Repositories
	deptRepo := postgres.NewDepartmentRepository(db)
	secRepo := postgres.NewSectionRepository(db)
	userRepo := postgres.NewUserRepository(db)
	tsRepo := postgres.NewTaskStatusRepository(db)
	taskRepo := postgres.NewTaskRepository(db)
	thRepo := postgres.NewTaskHistoryRepository(db)
	privRepo := postgres.NewPrivilegeRepository(db)

	// Services
	deptService := appDept.NewService(deptRepo)
	secService := appSec.NewService(secRepo)
	userService := appUser.NewService(userRepo)
	tsService := appTS.NewService(tsRepo)
	taskService := appTask.NewService(taskRepo)
	thService := appTH.NewService(thRepo)
	privService := appPriv.NewService(privRepo)
	authService := appAuth.NewService(userRepo, privRepo, jwtSecret, jwtExpiryHours)

	// Handlers
	deptHandler := handler.NewDepartmentHandler(deptService)
	secHandler := handler.NewSectionHandler(secService)
	userHandler := handler.NewUserHandler(userService)
	tsHandler := handler.NewTaskStatusHandler(tsService)
	taskHandler := handler.NewTaskHandler(taskService)
	thHandler := handler.NewTaskHistoryHandler(thService)
	privHandler := handler.NewPrivilegeHandler(privService)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")

	// Public routes (no authentication required)
	authHandler.RegisterRoutes(api)

	// Protected routes (authentication required)
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(authService))

	// Departments - protected with RBAC
	depts := protected.Group("/departments")
	depts.POST("", middleware.RequirePrivilege(authService, "departments", "create"), deptHandler.Create)
	depts.GET("", middleware.RequirePrivilege(authService, "departments", "read"), deptHandler.GetAll)
	depts.GET("/:id", middleware.RequirePrivilege(authService, "departments", "read"), deptHandler.GetByID)
	depts.PUT("/:id", middleware.RequirePrivilege(authService, "departments", "update"), deptHandler.Update)
	depts.DELETE("/:id", middleware.RequirePrivilege(authService, "departments", "delete"), deptHandler.Delete)

	// Sections - protected with RBAC
	sections := protected.Group("/sections")
	sections.POST("", middleware.RequirePrivilege(authService, "sections", "create"), secHandler.Create)
	sections.GET("", middleware.RequirePrivilege(authService, "sections", "read"), secHandler.GetAll)
	sections.GET("/:id", middleware.RequirePrivilege(authService, "sections", "read"), secHandler.GetByID)
	sections.PUT("/:id", middleware.RequirePrivilege(authService, "sections", "update"), secHandler.Update)
	sections.DELETE("/:id", middleware.RequirePrivilege(authService, "sections", "delete"), secHandler.Delete)

	// Users - protected with RBAC
	users := protected.Group("/users")
	users.POST("", middleware.RequirePrivilege(authService, "users", "create"), userHandler.Create)
	users.GET("", middleware.RequirePrivilege(authService, "users", "read"), userHandler.GetAll)
	users.GET("/:id", middleware.RequirePrivilege(authService, "users", "read"), userHandler.GetByID)
	users.PUT("/:id", middleware.RequirePrivilege(authService, "users", "update"), userHandler.Update)
	users.DELETE("/:id", middleware.RequirePrivilege(authService, "users", "delete"), userHandler.Delete)
	// User privileges management
	users.GET("/:id/privileges", middleware.RequirePrivilege(authService, "privileges", "read"), privHandler.GetUserPrivileges)
	users.POST("/:id/privileges", middleware.RequirePrivilege(authService, "privileges", "assign"), privHandler.AssignPrivilege)
	users.DELETE("/:id/privileges/:privilegeId", middleware.RequirePrivilege(authService, "privileges", "revoke"), privHandler.RevokePrivilege)

	// Task Statuses - protected with RBAC
	taskStatuses := protected.Group("/task-statuses")
	taskStatuses.POST("", middleware.RequirePrivilege(authService, "task-statuses", "create"), tsHandler.Create)
	taskStatuses.GET("", middleware.RequirePrivilege(authService, "task-statuses", "read"), tsHandler.GetAll)
	taskStatuses.GET("/:id", middleware.RequirePrivilege(authService, "task-statuses", "read"), tsHandler.GetByID)
	taskStatuses.PUT("/:id", middleware.RequirePrivilege(authService, "task-statuses", "update"), tsHandler.Update)
	taskStatuses.DELETE("/:id", middleware.RequirePrivilege(authService, "task-statuses", "delete"), tsHandler.Delete)

	// Tasks - protected with RBAC
	tasks := protected.Group("/tasks")
	tasks.POST("", middleware.RequirePrivilege(authService, "tasks", "create"), taskHandler.Create)
	tasks.GET("", middleware.RequirePrivilege(authService, "tasks", "read"), taskHandler.GetAll)
	tasks.GET("/:id", middleware.RequirePrivilege(authService, "tasks", "read"), taskHandler.GetByID)
	tasks.PUT("/:id", middleware.RequirePrivilege(authService, "tasks", "update"), taskHandler.Update)
	tasks.DELETE("/:id", middleware.RequirePrivilege(authService, "tasks", "delete"), taskHandler.Delete)

	// Task Histories - protected with RBAC
	taskHistories := protected.Group("/task-histories")
	taskHistories.POST("", middleware.RequirePrivilege(authService, "task-histories", "create"), thHandler.Create)
	taskHistories.GET("", middleware.RequirePrivilege(authService, "task-histories", "read"), thHandler.GetAll)
	taskHistories.GET("/:id", middleware.RequirePrivilege(authService, "task-histories", "read"), thHandler.GetByID)
	taskHistories.PUT("/:id", middleware.RequirePrivilege(authService, "task-histories", "update"), thHandler.Update)
	taskHistories.DELETE("/:id", middleware.RequirePrivilege(authService, "task-histories", "delete"), thHandler.Delete)

	// Privileges - protected with RBAC
	privileges := protected.Group("/privileges")
	privileges.POST("", middleware.RequirePrivilege(authService, "privileges", "create"), privHandler.Create)
	privileges.GET("", middleware.RequirePrivilege(authService, "privileges", "read"), privHandler.GetAll)
	privileges.GET("/:id", middleware.RequirePrivilege(authService, "privileges", "read"), privHandler.GetByID)
	privileges.PUT("/:id", middleware.RequirePrivilege(authService, "privileges", "update"), privHandler.Update)
	privileges.DELETE("/:id", middleware.RequirePrivilege(authService, "privileges", "delete"), privHandler.Delete)

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
