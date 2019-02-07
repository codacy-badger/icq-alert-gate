package icq

import (
	"net/url"
)

func (bot *Bot) SendIm(to, message string) (string, error) {
	params := url.Values{
		"t":       []string{to},
		"message": []string{message},
		//"mentions": []string{},
		//"parse": []string{}
	}
	data := ApiResponse{}
	err := bot.postText("/im/sendIM", params, &data)
	if err != nil {
		return "", err
	}
	return data.MsgId(), data.Error()
}
