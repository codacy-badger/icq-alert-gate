package web

import (
	"github.com/labstack/echo"
)

func (p *Provider) handlePUT(c echo.Context) error {
	data := c.Request()
	payload, err := p.payloadBySourceName(c.Param("source"))
	if err != nil {
		return err
	}
	messageString, err := payload.Parse(data)
	if err != nil {
		return err
	}
	_, err = p.Bot.SendIm(c.Param("target"), messageString)
	return err
}
