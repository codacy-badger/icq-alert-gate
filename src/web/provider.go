package web

import (
	"errors"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"icq"
	"webhook"
	"webhook/grafana"
	"webhook/jenkins"
	"webhook/raw"
)

var payloadSourceMap = map[string]webhook.Payload{
	"raw":              raw.Message{},
	"grafana":          grafana.Message{},
	"jenkins-outbound": jenkins.Message{},
}

// Provider represent single instances of bot and echo
type Provider struct {
	Bot      *icq.Bot
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
	e.PUT("/:source/:target", p.handleMessage)
	e.POST("/:source/:target", p.handleMessage)
	//
	p.instance = e

}

// Start prepare echo instance and start it
func (p *Provider) Start() error {
	p.initEcho()
	return p.instance.Start(":8888")
}
