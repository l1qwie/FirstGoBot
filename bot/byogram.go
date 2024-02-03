package bot

import (
	"firstgobot/byogram/formatter"
	"firstgobot/byogram/helper"
	"firstgobot/byogram/types"
	"fmt"
	"time"
)

func Receiving(tr *types.TelegramResponse) *formatter.Formatter {
	var (
		kbData      []string
		kbName      []string
		coordinates []int
		fm          formatter.Formatter
	)
	text := helper.ReturnText(tr)
	chatID := helper.ReturnChatId(tr)
	name := helper.ReturnName(tr)

	if text == "/start" {
		fm.WriteString(fmt.Sprintf("Hello, World! Hello, %s", name))
	} else if text == "/photo" {
		fm.AddPhotoFromTG("AgACAgQAAxkDAAIJRGW3rwaLqri1BkTdVQm1VFA8tE4HAAJeszEbEAABvFHW3MOANm9QFQEAAwIAA20AAzQE")
		fm.WriteString("Hello!")
	} else if text == "/keyboard" {
		kbName = []string{"I will send you a photo", "I will send you a video"}
		kbData = []string{"/photo", "/video"}
		coordinates = []int{1, 1}

		fm.SetIkbdDim(coordinates)
		for i := 0; i < len(kbName); i++ {
			fm.WriteInlineButtonCmd(kbName[i], kbData[i])
		}
		fm.WriteString(fmt.Sprintf("Hello! It's just a keyboard for a test, %s", name))

	} else if text == "/video" {
		fm.WriteString("Hello, firend!")
		fm.AddVideoFromTG("BAACAgIAAxkDAAIJW2W3sTguaruPGvo722qeKTcOPwvxAAIzPQACy-DASekiOEg76qGiNAQ")
	} else if text == "/wait" {
		time.Sleep(time.Second)
		fm.WriteString(fmt.Sprintf("Thank you for your wait, %s", name))
	} else {
		fm.WriteString(fmt.Sprintf("Sorry, I couldn't understand you, %s", name))
	}
	fm.WriteChatId(chatID)
	return &fm
}
