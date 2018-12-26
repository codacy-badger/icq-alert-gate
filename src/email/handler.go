package email

import (
	"bytes"
	"github.com/labstack/gommon/log"
	"io"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

const (
	bufferSize       int = 4 << 10
	icqMaxMessageLen int = 4 << 10
	icqPrefix            = "icq:"
)

func (p *Provider) handleConn(conn net.Conn) {
	err := conn.SetReadDeadline(time.Now().Add(time.Duration(5 * time.Second)))
	if err != nil {
		log.Warn("read connection timeout")
		return
	}
	data := make([]byte, bufferSize)
	_, err = io.ReadFull(conn, data)
	if err != nil {
		log.Warn(err)
	}
	err = p.handleData(&data)
	if err != nil {
		log.Warn(err)
	}
}

func (p *Provider) handleData(data *[]byte) error {
	var msg []byte
	var icqTargets, emailTargets []string
	message, err := mail.ReadMessage(bytes.NewReader(*data))
	if err != nil {
		return err
	}
	_, err = message.Body.Read(msg)
	if err != nil {
		return err
	}
	allTargets, err := mail.ParseAddressList(message.Header.Get("To"))
	if err != nil {
		return err
	}
	for _, target := range allTargets {
		if strings.HasPrefix(target.Address, icqPrefix) || strings.HasSuffix(target.Address, "@chat.agent") {
			icqTargets = append(icqTargets, strings.Replace(target.Address, icqPrefix, "", 1))
			continue
		}
		emailTargets = append(emailTargets, target.Address)
	}
	if len(emailTargets) > 0 {
		sender, err := mail.ParseAddress(message.Header.Get("From"))
		if err != nil {
			log.Warn(err)
		}
		log.Warn(p.forwardEmail(sender.Name, emailTargets, msg))
	}
	if len(icqTargets) > 0 {
		icqMessage := message.Header.Get("Subject") + "\n"
		if len(msg) > 0 {
			icqMessage += string(msg[:])
		}
		msgLen := len(icqMessage)
		if msgLen > icqMaxMessageLen {
			msgLen = icqMaxMessageLen
		}
		for _, target := range icqTargets {
			_, err := p.Bot.SendIm(target, icqMessage[:msgLen])
			if err != nil {
				log.Warn(err)
			}
		}
	}
	return nil
}

func (p *Provider) forwardEmail(identity string, targets []string, message []byte) error {
	auth := smtp.PlainAuth(identity, p.Username, p.Password, p.Host)
	if p.Password != "" {
		return smtp.SendMail(p.Host, auth, p.Username, targets, message)
	}
	return smtp.SendMail(p.Host, nil, p.Username, targets, message)

}
