package routers

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"github.com/liukeshao/echo-admin/pkg/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/liukeshao/echo-admin/docs"
)

// BuildRouter builds the router
func (e *Endpoints) BuildRouter() error {
	// default route group
	g := e.container.Web.Group("")
	g.Use(
		echomw.Recover(),
		echomw.RequestID(),
		middleware.SetLogger(),
		middleware.LogRequest(),
		echomw.Gzip(),
		echomw.TimeoutWithConfig(
			echomw.TimeoutConfig{
				Timeout: e.container.Config.App.Timeout,
			},
		),
	)
	e.container.Web.HTTPErrorHandler = customHTTPErrorHandler

	// swagger
	g.GET("/swagger/*", echoSwagger.WrapHandler)

	jwtmw := echojwt.WithConfig(echojwt.Config{SigningKey: []byte(e.container.Config.App.EncryptionKey)})
	strictAuth := g.Group("", jwtmw)
	noAuth := g.Group("", middleware.RequireNoAuthentication())

	accountNoAuthGroup := noAuth.Group("/account")
	{
		accountNoAuthGroup.POST("/login", e.account.Login).Name = "login"
		accountNoAuthGroup.POST("/register", e.account.Register).Name = "register"
	}

	accountGroup := strictAuth.Group("/account")
	{
		accountGroup.GET("/profile", e.account.Profile).Name = "profile"
		accountGroup.GET("/logout", e.account.Logout).Name = "logout"
	}

	// pipeline
	pipelineGroup := strictAuth.Group("/api/pipelines")
	{
		pipelineGroup.POST("/:pipelineID/actions/run", nil)
	}

	return nil
}
