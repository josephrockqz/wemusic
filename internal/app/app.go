package app

import (
	"github.com/josephrockqz/wemusic-golang/internal/config"
	"github.com/josephrockqz/wemusic-golang/internal/services"
	"github.com/josephrockqz/wemusic-golang/internal/transport/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

func Run() {
	e := echo.New()

	e.Logger.SetLevel(log.INFO)
	zap.ReplaceGlobals(zap.Must(config.CreateLogger()))

	e.Use(middleware.CorsMiddleware())

	e.GET("/spotify-login", services.SpotifyLogin)
	e.GET("/spotify-user-authorization-callback", services.SpotifyUserAuthorizationCallback)

	e.Logger.Fatal(e.Start(":8080"))
}
