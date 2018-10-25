package app

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io"
	"io/ioutil"
)

func (p *Provider) handleMessage(c echo.Context) error {
	target := c.Param("target")
	data := c.Request().Body
	defer data.Close()
	message, err := parseMessage(data)
	if err != nil {
		return err
	}
	_, err = p.Bot.SendIm(target, message)
	return err
}

func parseMessage(data io.ReadCloser) (string, error) {
	messageBytes, err := ioutil.ReadAll(data)
	if err != nil {
		return "", err
	}
	grafanaMsg := GrafanaMessage{}
	err = json.Unmarshal(messageBytes, &grafanaMsg)
	if err != nil {
		return "", err
	}
	return grafanaMsg.String(), nil
}
