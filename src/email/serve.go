package email

import (
	"github.com/labstack/gommon/log"
	"net"
)

func (p *Provider) Serve(address string) error {
	listner, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer listner.Close()
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Warn(err)
			continue
		}
		go p.handleConn(conn)
	}
}
