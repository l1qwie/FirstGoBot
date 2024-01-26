package bot

import (
	redirection "firstgobot/bot/redirection"
	by "firstgobot/byogram/methods"
)

func doSend(text, image, keyboard string, chatID int) {
	if text != "" {
		if keyboard != "" {
			by.SendMessageWithKeyboard(text, keyboard, chatID)
		} else {
			by.SendMessage(text, chatID)
		}
	} else if image != "" {
		by.SendPhoto(image, chatID)
	}
}

func Acceptance(phrase, name string, chatID int) {
	var (
		text     string
		image    string
		keyboard string
	)
	text, image, keyboard = redirection.DispatcherPhrase(phrase, name, chatID)

	doSend(text, image, keyboard, chatID)
}
