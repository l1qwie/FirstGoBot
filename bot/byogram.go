package bot

import (
	redirection "firstgobot/bot/redirection"
	"firstgobot/byogram/formatter"
	"log"
)

func doSend(text, image, video string, kbName, kbData []string, chatID int) {
	var (
		fm          formatter.Formatter
		coordinates []int
		err         error
	)

	if text != "" {
		if kbName != nil && kbData != nil {
			fm.WriteString(text)
			fm.WriteChatId(chatID)
			for i := 0; i < 2; i++ {
				coordinates = append(coordinates, 1)
			}
			fm.SetIkbdDim(coordinates)
			for i := 0; i < 2; i++ {
				fm.WriteInlineButtonCmd(kbName[i], kbData[i])
			}
			//fm.WriteInlineButtonCmd(kbName, kbData)
			//fm.WriteInlineButtonCmd(kbName, kbData)
			//fm.WriteInlineButtonUrl(kbName, kbData)
			err = fm.SendMessage()
		} else if image != "" {
			fm.AddPhotoFromTG(image)
			fm.WriteString(text)
			fm.WriteChatId(chatID)
			err = fm.SendMessage()
		} else if video != "" {
			fm.AddVideoFromTG(video)
			fm.WriteString(text)
			fm.WriteChatId(chatID)
			err = fm.SendMessage()
		} else {
			fm.WriteString(text)
			fm.WriteChatId(chatID)
			err = fm.SendMessage()
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}

func Acceptance(phrase, name string, chatID int) {
	var (
		text, image, video string
		kbName, kbData     []string
	)
	text, image, video, kbName, kbData = redirection.DispatcherPhrase(phrase, name, chatID)

	doSend(text, image, video, kbName, kbData, chatID)
}
