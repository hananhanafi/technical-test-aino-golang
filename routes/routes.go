package routes

import (
	"github.com/hananhanafi/technical-test-aino-golang/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	IsAuthenticated := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	})

	// user group
	userAPI := e.Group("/api/user")
	// user API routes
	userAPI.GET("/get-all", controllers.FetchAllUser, IsAuthenticated)
	userAPI.POST("/register", controllers.Register)
	userAPI.POST("/login", controllers.Login)

	return e
}
