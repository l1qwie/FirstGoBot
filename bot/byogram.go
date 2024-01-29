package bot

import (
	redirection "firstgobot/bot/redirection"
	"firstgobot/byogram/formatter"
)

func doSend(text, image, kbName, kbData string, chatID int) {
	var (
		fm          formatter.Formatter
		coordinates []int
	)

	if text != "" {
		if kbName != "" && kbData != "" {
			fm.WriteString(text)
			fm.WriteChatId(chatID)
			for i := 0; i < 2; i++ {
				coordinates = append(coordinates, 1)
			}
			fm.SetIkbdDim(coordinates)
			fm.WriteInlineButtonCmd(kbName, kbData)
			fm.WriteInlineButtonCmd(kbName, kbData)
			//fm.WriteInlineButtonUrl(kbName, kbData)
			fm.SendMessage()
		} else if image != "" {
			//methods.SendPhoto(image, chatID)
			fm.AddPhotoFromMemmory(image)
			//fm.WriteString(text)
			fm.WriteChatId(chatID)
			fm.SendMessage()
		} else {
			fm.WriteString(text)
			fm.WriteChatId(chatID)
			fm.SendMessage()
		}
	}
}

func Acceptance(phrase, name string, chatID int) {
	var (
		text, image, kbName, kbData string
	)
	text, image, kbName, kbData = redirection.DispatcherPhrase(phrase, name, chatID)

	doSend(text, image, kbName, kbData, chatID)
}
