package bot

import (
	"firstgobot/byogram/types"
	"fmt"
)

func Receiving(tr types.TelegramResponse, fm types.FMTRS) {
	var (
		kbData      []string
		kbName      []string
		coordinates []int
	)

	if tr.Result[0].Message.Text == "/start" || tr.Result[0].Query.Data == "/start" {
		fm.WriteString(fmt.Sprintf("Hello, World! Hello, %s", tr.Result[0].Message.TypeFrom.Name))
	} else if tr.Result[0].Message.Text == "/photo" || tr.Result[0].Query.Data == "/photo" {
		fm.AddPhotoFromTG("AgACAgQAAxkDAAIJRGW3rwaLqri1BkTdVQm1VFA8tE4HAAJeszEbEAABvFHW3MOANm9QFQEAAwIAA20AAzQE")
		fm.WriteString("Hello!")
	} else if tr.Result[0].Message.Text == "/keyboard" || tr.Result[0].Query.Data == "/keyboard" {
		kbName = []string{"I will send you a photo", "I will send you a video"}
		kbData = []string{"/photo", "/video"}
		coordinates = []int{1, 1}

		fm.SetIkbdDim(coordinates)
		for i := 0; i < len(kbName); i++ {
			fm.WriteInlineButtonCmd(kbName[i], kbData[i])
		}
		fm.WriteString(fmt.Sprintf("Hello! It's just a keyboard for a test, %s", tr.Result[0].Message.TypeFrom.Name))

	} else if tr.Result[0].Message.Text == "/video" || tr.Result[0].Query.Data == "/video" {
		fm.WriteString("Hello, firend!")
		fm.AddVideoFromTG("BAACAgIAAxkDAAIJW2W3sTguaruPGvo722qeKTcOPwvxAAIzPQACy-DASekiOEg76qGiNAQ")
	} else {
		fm.WriteString(fmt.Sprintf("Sorry, I couldn't understand you, %s", tr.Result[0].Message.TypeFrom.Name))
	}
	fm.WriteChatId(tr.Result[0].Message.TypeFrom.UserID)
}
