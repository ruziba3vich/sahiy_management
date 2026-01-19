package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/sahiy_management/internal/application"
	"github.com/ruziba3vich/sahiy_management/internal/application/auth"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/handler"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ruziba3vich/sahiy_management/docs"
)

func SetupRoutes(r *gin.Engine, svc *application.Service) {
	// Handlers
	deptHandler := handler.NewDepartmentHandler(svc.GetDepartment())
	secHandler := handler.NewSectionHandler(svc.GetSection())
	userHandler := handler.NewUserHandler(svc.GetUser())
	tsHandler := handler.NewTaskStatusHandler(svc.GetTaskStatus())
	taskHandler := handler.NewTaskHandler(svc.GetTask())
	thHandler := handler.NewTaskHistoryHandler(svc.GetTaskHistory())
	privHandler := handler.NewPrivilegeHandler(svc.GetPrivilages())
	authHandler := handler.NewAuthHandler(svc.GetAuth())

	authService := svc.GetAuth()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")

	// Public routes
	authHandler.RegisterRoutes(api)

	// Protected routes
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(authService))

	// Departments
	registerCRUD(protected, "/departments", "departments", authService, deptHandler)

	// Sections
	registerCRUD(protected, "/sections", "sections", authService, secHandler)

	// Users
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

	// Task Statuses
	registerCRUD(protected, "/task-statuses", "task-statuses", authService, tsHandler)

	// Tasks
	registerCRUD(protected, "/tasks", "tasks", authService, taskHandler)

	// Task Histories
	registerCRUD(protected, "/task-histories", "task-histories", authService, thHandler)

	// Privileges
	registerCRUD(protected, "/privileges", "privileges", authService, privHandler)
}

type crudHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func registerCRUD(rg *gin.RouterGroup, path, resource string, authService *auth.Service, h crudHandler) {
	group := rg.Group(path)
	group.POST("", middleware.RequirePrivilege(authService, resource, "create"), h.Create)
	group.GET("", middleware.RequirePrivilege(authService, resource, "read"), h.GetAll)
	group.GET("/:id", middleware.RequirePrivilege(authService, resource, "read"), h.GetByID)
	group.PUT("/:id", middleware.RequirePrivilege(authService, resource, "update"), h.Update)
	group.DELETE("/:id", middleware.RequirePrivilege(authService, resource, "delete"), h.Delete)
}
