package app

import (
	"github.com/josephrockqz/wemusic-golang/internal/services"
	"github.com/josephrockqz/wemusic-golang/internal/transport/middleware"
	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	e.Use(middleware.LoggerMiddleWare())
	e.Use(middleware.CorsMiddleware())

	e.GET("/spotify-login", services.SpotifyLogin)
	e.GET("/spotify-user-authorization-callback", services.SpotifyUserAuthorizationCallback)

	e.Logger.Fatal(e.Start(":8080"))
}
