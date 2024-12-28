package routes

import (
	"myproject/controllers"
	"myproject/middleware"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	e.GET("/", controllers.Homepage)
	e.GET("/register", controllers.Register)
	e.GET("/login", controllers.Login)
	e.POST("/registerUser", controllers.Register_user)
	e.POST("/loginHere", controllers.LoginHere)
	e.GET("/user_dashboard", controllers.User_dashboard, middleware.LoginMiddleware, middleware.UserMiddleware)
	e.GET("/admin_dashboard", controllers.Admin_dashboard, middleware.LoginMiddleware, middleware.AdminMiddleware)
	e.POST("/createtask", controllers.CreateTask, middleware.LoginMiddleware, middleware.UserMiddleware)
	e.GET("/logout", controllers.Logout, middleware.LoginMiddleware)
	e.POST("/updatetask/:id", controllers.UpdateTask, middleware.LoginMiddleware, middleware.UserMiddleware)
	e.POST("/deletetask/:id", controllers.DeleteTask, middleware.LoginMiddleware, middleware.UserMiddleware)
}
