package icq

func (bot *Bot) SendFile(to, filePath string) (string, error) {
	fileUrl, err := bot.UploadFile(to, filePath)
	if err != nil {
		return "", err
	}
	return bot.SendIm(to, fileUrl)
}
