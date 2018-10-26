package app

import (
	"alert"
	"alert/grafana"
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/slavyan85/gocq"
)

var payloadSourceMap = map[string]alert.Payload{
	"grafana": grafana.GrafanaMessage{},
}

type Provider struct {
	Bot      *gocq.Bot
	instance *echo.Echo
}

func (Provider) payloadBySourceName(sourceName string) (alert.Payload, error) {
	payload, ok := payloadSourceMap[sourceName]
	if !ok {
		return nil, errors.New("unknown alert source")
	}
	return payload, nil
}

func (p *Provider) initEcho() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//
	e.PUT("/:source/:target", p.handlePUT)
	//
	p.instance = e

}

func (p *Provider) Start() {
	p.initEcho()
	p.instance.Logger.Fatal(p.instance.Start(":8888"))
}
