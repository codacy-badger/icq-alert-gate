package web

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/slavyan85/gocq"
	"webhook"
	"webhook/grafana"
)

var payloadSourceMap = map[string]webhook.Payload{
	"grafana": grafana.GrafanaMessage{},
}

type Provider struct {
	Bot      *gocq.Bot
	instance *echo.Echo
}

func (Provider) payloadBySourceName(sourceName string) (webhook.Payload, error) {
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

func (p *Provider) Start(failChan chan error) {
	p.initEcho()
	failChan <- p.instance.Start(":8888")
}
