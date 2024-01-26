package redirection

import (
	"firstgobot/byogram/types"
	"fmt"
)

func DispatcherPhrase(phrase, name string, chatID int) (string, string, string) {
	var (
		text     string
		image    string
		KB       string
		keyboard types.InlineKeyboardMarkup
	)

	if phrase == "/start" {
		text = fmt.Sprintf("Hello, World! Hello, %s", name)
	} else if phrase == "/photo" {
		image = "FOOTBALL1.jpg"
	} else if phrase == "/keyboard" {
		text = fmt.Sprintf("Hello! It's just a keyboard for a test, %s", name)
		keyboard = types.InlineKeyboardMarkup{
			Buttons: [][]types.InlineKeyboardButton{
				{
					{Text: "Нажми меня", CallbackData: "button_pressed"},
					{Text: "Нажми меня еще раз", CallbackData: "/start"},
				},
			},
		}
		KB = types.CreateInlineKeyoard(keyboard)
	} else {
		text = fmt.Sprintf("Sorry, I couldn't understand you, %s", name)
	}

	return text, image, KB
}
