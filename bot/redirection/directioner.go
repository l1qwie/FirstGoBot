package redirection

import (
	"fmt"
)

func DispatcherPhrase(phrase, name string, chatID int) (string, string, string, []string, []string) {
	var (
		text, image, video string
		kbData, kbName     []string
	)

	if phrase == "/start" {
		text = fmt.Sprintf("Hello, World! Hello, %s", name)
	} else if phrase == "/photo" {
		image = "AgACAgQAAxkDAAIJRGW3rwaLqri1BkTdVQm1VFA8tE4HAAJeszEbEAABvFHW3MOANm9QFQEAAwIAA20AAzQE"
		text = "Hello!"
	} else if phrase == "/keyboard" {
		text = fmt.Sprintf("Hello! It's just a keyboard for a test, %s", name)
		kbName = append(kbName, "I will send you a photo")
		kbName = append(kbName, "I will send you a video")
		kbData = append(kbData, "/photo")
		kbData = append(kbData, "/video")
	} else if phrase == "/video" {
		text = "Shalom!"
		video = "BAACAgIAAxkDAAIJW2W3sTguaruPGvo722qeKTcOPwvxAAIzPQACy-DASekiOEg76qGiNAQ"
	} else {
		text = fmt.Sprintf("Sorry, I couldn't understand you, %s", name)
	}

	return text, image, video, kbName, kbData
}
