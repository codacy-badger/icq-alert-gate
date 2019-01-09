package web

import (
	"github.com/labstack/echo"
)

func (p *Provider) handleMessage(c echo.Context) error {
	payload, err := p.payloadBySourceName(c.Param("source"))
	if err != nil {
		return err
	}
	messageString, err := payload.Parse(c.Request())
	if err != nil {
		return err
	}
	_, err = p.Bot.SendIm(c.Param("target"), messageString)
	return err
}
