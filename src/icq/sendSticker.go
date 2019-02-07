package icq

import (
	"net/url"
)

func (bot *Bot) SendSticker(to, stickerId string) error {
	params := url.Values{
		"t":         []string{to},
		"stickerId": []string{stickerId},
	}
	data := ApiResponse{}
	err := bot.postText("/im/sendSticker", params, &data)
	if err != nil {
		return err
	}
	return data.Error()
}
