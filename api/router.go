package api

import (
	"fmt"
	"time"

	"Trial/BANK-NOVANNA/api/middleware"
	"Trial/BANK-NOVANNA/internal/domain/handler"
	"Trial/BANK-NOVANNA/internal/domain/repository"
	"Trial/BANK-NOVANNA/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	custRepo := repository.NewCustomerRepository(db)
	custService := service.NewCustomerService(custRepo)
	custHandler := handler.NewCustomerHandler(custService)

	customer := router.Group("v1/customer")

	customer.POST("/", middleware.AuthMiddleware(), custHandler.SaveCustomer)
	customer.GET("/:no_ktp", middleware.AuthMiddleware(), custHandler.GetDetailCustomer)
	customer.GET("/", middleware.AuthMiddleware(), custHandler.GetAllCustomer)
	customer.PUT("/:cust_id", middleware.AuthMiddleware(), custHandler.UpdateCustomer)
	customer.DELETE("/:no_ktp", middleware.AuthMiddleware(), custHandler.DeleteCustomer)

	//Register User Repo
	adminRepo := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)

	admin := router.Group("v1/admin")

	admin.POST("/", adminHandler.RegisterAdmin)
	admin.POST("/login", adminHandler.Login)

	return router
}
