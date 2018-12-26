package email

import (
	"github.com/slavyan85/gocq"
)

type Provider struct {
	Bot      *gocq.Bot
	Username string
	Password string
	Host     string
}

func (p *Provider) Start(failChan chan error) {
	if p.Host == "" {
		return
	}
	failChan <- p.Serve(":2525")
}
