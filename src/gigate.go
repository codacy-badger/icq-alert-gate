package main

import (
	"app"
	"github.com/labstack/gommon/log"
	"github.com/slavyan85/gocq"
	"os"
)

func init() {
	log.SetLevel(log.INFO)
}

func main() {
	botUin := os.Getenv("BOT_UIN")         // required
	botToken := os.Getenv("BOT_TOKEN")     // required
	botNick := os.Getenv("BOT_NICK")       // required
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
