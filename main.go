package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vtomkiv/golang.api/middleware/logg"
	"github.com/vtomkiv/golang.api/middleware/auth/fb"
)


func main() {
	e := echo.New()

	//init logger
	e.Use(logg.Logger())

	//login route
	e.GET("/login", fb.Login )

	//fb auth callback
	e.GET("/auth/fb/callback", fb.HandleFBCallback)

	// Restricted group
	r := e.Group("/authorized")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &fb.JwtFBClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.GET("/ping", fb.Restricted)


	e.Logger.Fatal(e.Start(":8088"))
}

