package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/sahiy_management/internal/application"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/handler"
	"github.com/ruziba3vich/sahiy_management/internal/interface/http/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ruziba3vich/sahiy_management/docs"
)

func SetupRoutes(r *gin.Engine, svc *application.Service) {
	// Handlers
	deptHandler := handler.NewDepartmentHandler(svc.GetDepartment())
	scheduleHandler := handler.NewScheduleHandler(svc.GetSchedule())
	userHandler := handler.NewUserHandler(svc.GetUser())
	userActionHandler := handler.NewUserActionHandler(svc.GetUserAction())
	tsHandler := handler.NewTaskStatusHandler(svc.GetTaskStatus())
	taskHandler := handler.NewTaskHandler(svc.GetTask())
	thHandler := handler.NewTaskHistoryHandler(svc.GetTaskHistory())
	branchHandler := handler.NewBranchHandler(svc.GetBranch())
	sectionHandler := handler.NewSectionHandler(svc.GetSection())
	authHandler := handler.NewAuthHandler(svc.GetAuth())
	roleHandler := handler.NewRoleHandler()
	statsHandler := handler.NewStatisticsHandler(svc.GetStatistics())

	authService := svc.GetAuth()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")

	// Public routes
	authHandler.RegisterRoutes(api)

	// Protected routes (authentication only, no authorization)
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(authService))

	// Departments
	deptHandler.RegisterRoutes(protected)

	// Branches
	branchHandler.RegisterRoutes(protected)

	// Sections
	sectionHandler.RegisterRoutes(protected)

	// Schedules
	registerCRUD(protected, "/schedules", scheduleHandler)

	// Users
	registerCRUD(protected, "/users", userHandler)

	// User Actions
	userActionHandler.RegisterRoutes(protected)

	// Task Statuses
	registerCRUD(protected, "/task-statuses", tsHandler)
	protected.GET("/task-statuses/types", tsHandler.GetTypes)

	// Tasks
	registerCRUD(protected, "/tasks", taskHandler)
	protected.GET("/tasks/my", taskHandler.GetMyTasks)
	protected.GET("/tasks/calendar", taskHandler.GetCalendarEvents)
	protected.GET("/tasks/calendar/:id", taskHandler.GetCalendarEventByID)

	// Task Histories
	registerCRUD(protected, "/task-histories", thHandler)
	protected.GET("/task-histories/my", thHandler.GetMyHistory)

	// Roles
	protected.GET("/roles", roleHandler.GetAll)

	// Statistics
	statsHandler.RegisterRoutes(protected)
}

type crudHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func registerCRUD(rg *gin.RouterGroup, path string, h crudHandler) {
	group := rg.Group(path)
	group.POST("", h.Create)
	group.GET("", h.GetAll)
	group.GET("/:id", h.GetByID)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}
