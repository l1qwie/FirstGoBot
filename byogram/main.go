package byogram

import (
	"firstgobot/bot"
	"firstgobot/byogram/executer"
	"firstgobot/byogram/formatter"
	"firstgobot/byogram/tests"
	"firstgobot/byogram/types"
	"fmt"
)

func Reset(telegramResponse *types.TelegramResponse) {
	*telegramResponse = types.TelegramResponse{
		Ok: false,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 0,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   0,
						IsBot:    false,
						Name:     "",
						LastName: "",
						Username: "",
					},
					Text: "",
				},
				Query: types.Callback{
					TypeFrom: types.User{
						UserID:   0,
						IsBot:    false,
						Name:     "",
						LastName: "",
						Username: "",
					},
					Data: "",
				},
			},
		},
	}
}

func StartWithTelegram() {
	var (
		offset           int
		fm               formatter.Formatter
		err              error
		telegramResponse types.TelegramResponse
	)

	err = executer.RequestOffset(types.TelebotToken, &offset)
	for err != nil {
		err = executer.RequestOffset(types.TelebotToken, &offset)
	}
	for {
		err = executer.Updates(types.TelebotToken, &offset, &telegramResponse)
		if err != nil {
			fmt.Println("Обновления не были получены: ", err)
		} else if len(telegramResponse.Result) == 0 {
		} else {
			fmt.Println(telegramResponse)
			fmt.Println(offset)
			bot.Receiving(telegramResponse, &fm)
			fm.Send()
			Reset(&telegramResponse)
			offset = offset + 1
		}
	}
}

func StartTests() {
	var (
		//err              error
		tfm              tests.Formatter
		telegramResponse types.TelegramResponse
	)
	telegramResponse = types.TelegramResponse{
		Ok: true,
		Result: []types.TelegramUpdate{
			{
				UpdateID: 123,
				Message: types.InfMessage{
					TypeFrom: types.User{
						UserID:   456,
						IsBot:    false,
						Name:     "John",
						LastName: "Doe",
						Username: "johndoe",
					},
					Text: "/keyboard",
				},
			},
		},
	}

	bot.Receiving(telegramResponse, &tfm)
	_ = tfm.AssertString(fmt.Sprintf("Hello! It's just a keyboard for a test, %s", telegramResponse.Result[0].Message.TypeFrom.Name), true)
	_ = tfm.AssertChatId(456, true)
	//[]int{1, 1} []string{"I will send you a photo", "I will send you a video"} []string{"/photo", "/video"} []string{"cmd", "cmd"}
	_ = tfm.AssertInlineKeyboard([]int{1, 1}, []string{"I will send you a photo", "I will send you a video"}, []string{"/photo", "/video"}, []string{"cmd", "cmd"}, true)
	//_ = tfm.AssertPhoto("AgACAgQAAxkDAAIJRGW3rwaLqri1BkTdVQm1VFA8tE4HAAJeszEbEAABvFHW3MOANm9QFQEAAwIAA20AAzQE", true)
	//_ = tfm.AssertVideo("BAACAgIAAxkDAAIJW2W3sTguaruPGvo722qeKTcOPwvxAAIzPQACy-DASekiOEg76qGiNAQ", true)
	fmt.Print("All was alright!")

}
