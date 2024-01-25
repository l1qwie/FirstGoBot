package bot

import (
	redirection "firstgobot/bot/redirection"
	by "firstgobot/byogram/methods"
)

func doSend(text, image string, chatID int) {
	if text != "" {
		by.SendMessage(text, chatID)
	} else if image != "" {
		by.SendPhoto(image, chatID)
	}
}

func Acceptance(phrase, name string, chatID int) {
	var (
		text  string
		image string
	)
	text, image = redirection.DispatcherPhrase(phrase, name, chatID)

	doSend(text, image, chatID)
}
