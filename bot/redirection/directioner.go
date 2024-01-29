package redirection

import (
	"fmt"
)

func DispatcherPhrase(phrase, name string, chatID int) (string, string, string, string) {
	var (
		text, image, kbName, kbData string
	)

	if phrase == "/start" {
		text = fmt.Sprintf("Hello, World! Hello, %s", name)
	} else if phrase == "/photo" {
		image = "FOOTBALL1.jpg"
		text = "Hello!"
	} else if phrase == "/keyboard" {
		text = fmt.Sprintf("Hello! It's just a keyboard for a test, %s", name)
		kbName = "Press On Me!"
		kbData = "https://core.telegram.org/"
	} else {
		text = fmt.Sprintf("Sorry, I couldn't understand you, %s", name)
	}

	return text, image, kbName, kbData
}
