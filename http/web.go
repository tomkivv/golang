package http

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vtomkiv/golang.api/http/handler"
)

type ControllerContext struct {
	TaskController handler.TaskController
} 

func (cc ControllerContext) Run()  {
	e := echo.New()

	//init logger
	e.Use(middleware.Logger())

	//login route
	e.GET("/login", handler.FBLogin )

	//fb auth callback
	e.GET("/auth/fb/callback", handler.HandleFBCallback)

	// Restricted group
	r := e.Group("/authorized")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     handler.JwtFBClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	// Routes
	e.POST("/tasks", cc.TaskController.CreateTask)
	e.GET("/tasks/:id", cc.TaskController.FindTask)
	e.PUT("/tasks", cc.TaskController.UpdateTask)
	e.DELETE("/tasks/:id", cc.TaskController.DeleteTask)


	e.Logger.Fatal(e.Start(":8088"))

}
