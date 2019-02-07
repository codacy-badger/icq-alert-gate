package icq

func (bot *Bot) UploadFile(to, filePath string) (string, error) {
	response := struct {
		Data struct {
			StaticUrl string `json:"static_url"`
		} `json:"data"`
	}{}
	err := bot.postFile(filePath, &response)
	if err != nil {
		return "", err
	}
	return response.Data.StaticUrl, nil
}
