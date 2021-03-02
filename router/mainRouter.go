package router

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"

	. "kamgo/config"
	"kamgo/modules/log"
	"kamgo/router/api"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func InitRoutes() map[string]*Host {
	hosts := make(map[string]*Host)
	return hosts
}

func RunSubdomains(confFilePath string) {
	if err := InitConfig(confFilePath); err != nil {
		log.Panic(err)
	}

	log.SetLevel(GetLogLvl())

	// Server
	e := echo.New()

	e.Pre(mw.RemoveTrailingSlash())

	e.Logger.SetLevel(GetLogLvl())

	// Secure, XSS/CSS HSTS
	e.Use(mw.SecureWithConfig(mw.DefaultSecureConfig))
	e.Use(mw.MethodOverride())

	hosts := InitRoutes()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		u, _err := url.Parse(c.Scheme() + "://" + req.Host)
		if _err != nil {
			e.Logger.Errorf("Request URL parse error:%v", _err)
		}

		hosts[u.Hostname()] = &Host{api.Routers()}
		hosts[u.Hostname()].Echo.ServeHTTP(res, req)
		return
	})

	if !Conf.Server.Graceful {
		e.Logger.Fatal(e.Start(Conf.Server.Addr))
	} else {
		// Graceful Shutdown
		// Start server
		go func() {
			if err := e.Start(Conf.Server.Addr); err != nil {
				e.Logger.Errorf("Shutting down the server with error:%v", err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}
}
