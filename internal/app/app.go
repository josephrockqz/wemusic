package app

import (
	"github.com/josephrockqz/wemusic-golang/internal/services"
	"github.com/josephrockqz/wemusic-golang/internal/transport/middleware"
	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	// TODO: use additional middleware (e.g. logger)
	e.Use(middleware.CorsMiddleware())

	e.GET("/spotify-login", services.SpotifyLogin)

	e.Start(":8080")
}
