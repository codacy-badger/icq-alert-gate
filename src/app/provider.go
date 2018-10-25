package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/slavyan85/gocq"
)

type Provider struct {
	Bot      *gocq.Bot
	instance *echo.Echo
}

func (p *Provider) initEcho() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//
	e.PUT("/:target", p.handleMessage)
	//
	p.instance = e

}

func (p *Provider) Start() {
	p.initEcho()
	p.instance.Logger.Fatal(p.instance.Start(":8888"))
}
