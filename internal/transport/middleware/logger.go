package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		/*
			fields removed from default config:

				"remote_ip":"${remote_ip}"
				"user_agent":"${user_agent}"
				"latency":${latency}
				"latency_human":"${latency_human}"
				"bytes_in":${bytes_in}
				"bytes_out":${bytes_out}
		*/
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}",` +
			`"status":${status},"error":"${error}}"` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	})
}
