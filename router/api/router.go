package api

import (
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	. "kamgo/config"
)

func Routers() *echo.Echo {
	e := echo.New()
	e.Use(NewContext())
	if Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("api")
	e.Logger.SetLevel(GetLogLvl())

	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Validator = NewValidator()

	e.Use(mw.RequestIDWithConfig(mw.RequestIDConfig{
		Generator: func() string {
			return genUUID()
		},
	}))

	e.POST("v1/Subscribers", handler(CreateSipEndpoint))
	e.DELETE("v1/Subscribers/:endpoint", handler(DeleteSipEndpoint))
	e.GET("v1/Subscribers/:endpoint", handler(GetSipEndpoint))
	e.GET("v1/Subscribers", handler(GetListSipEndpoint))
	e.POST("v1/Subscribers/:endpoint", handler(UpdateSipEndpoint))
	e.PUT("v1/Subscribers/:endpoint", handler(UpdateSipEndpoint))
	return e
}

type (
	HandlerFunc func(*Context) error
)

func handler(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*Context)
		return h(ctx)
	}
}
func genUUID() string {
	requestUUID, err := uuid.NewV4()
	if err != nil {
		//add error handler
	}
	return requestUUID.String()
}
