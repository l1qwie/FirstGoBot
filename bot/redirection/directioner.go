package redirection

import "fmt"

func DispatcherPhrase(phrase, name string, chatID int) (string, string) {
	var (
		text  string
		image string
	)

	if phrase == "/start" {
		text = fmt.Sprint("Hello, World! Hello, ", name)
	} else if phrase == "/photo" {
		image = "FOOTBALL1.jpg"
	} else {
		text = fmt.Sprint("Sorry, I couldn't understand you, ", name)
	}

	return text, image
}
