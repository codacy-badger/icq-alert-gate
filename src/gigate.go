package main

import (
	"app"
	"github.com/labstack/gommon/log"
	"github.com/slavyan85/gocq"
	"os"
)

func init() {
	log.SetLevel(log.INFO)
	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	botToken := os.Getenv("BOT_TOKEN")     // required
	botUin := os.Getenv("BOT_UIN")         // optional
	botNick := os.Getenv("BOT_NICK")       // optional
	botName := os.Getenv("BOT_NAME")       // optional
	botVersion := os.Getenv("BOT_VERSION") // optional
	if botName == "" {
		botName = botNick
	}
	bot := gocq.Bot{
		//ApiURL:"http://localhost:5000",
		Name:    botName,
		Nick:    botNick,
		Uin:     botUin,
		Token:   botToken,
		Version: botVersion,
	}
	err := bot.Init()
	if err != nil {
		log.Fatal(err)
	}
	provider := app.Provider{Bot: &bot}
	provider.Start()
}
