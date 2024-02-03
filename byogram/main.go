package byogram

import (
	"firstgobot/bot"
	"firstgobot/byogram/executer"
	"firstgobot/byogram/formatter"
	"firstgobot/byogram/types"
	"fmt"
	"log"
	"time"
)

func pollResponse(responses chan<- *types.TelegramResponse) {
	var (
		offset           int
		err              error
		telegramResponse *types.TelegramResponse
	)

	err = executer.RequestOffset(types.TelebotToken, &offset)
	for err != nil {
		err = executer.RequestOffset(types.TelebotToken, &offset)
	}
	for {
		telegramResponse = new(types.TelegramResponse)
		err = executer.Updates(types.TelebotToken, &offset, telegramResponse)
		if err != nil {
			log.Fatal(err)
		}
		if len(telegramResponse.Result) != 0 {
			responses <- telegramResponse
			offset = offset + 1
		}
	}

	/*for response := range responses {
		fmt.Println(offset)
		telegramResponse = new(types.TelegramResponse)
		if err != nil {
			fmt.Println("Обновления не были получены: ", err)
		} else if len(telegramResponse.Result) == 0 {
		} else {
			fmt.Println(telegramResponse)
			fmt.Println(offset)
			go worker(*telegramResponse, newFm())
			offset = offset + 1
		}
	}*/
}
func worker(r *types.TelegramResponse, output chan<- *formatter.Formatter) {
	fm := bot.Receiving(r)
	fm.Complete()
	output <- fm
}

func pushRequest(requests <-chan *formatter.Formatter) {
	for r := range requests {
		err := r.Send()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func dispatcher(input <-chan *types.TelegramResponse, output chan<- *formatter.Formatter) {
	for r := range input {
		go worker(r, output)
	}
}

func StartWithTelegram() {
	var responses chan *types.TelegramResponse
	var requests chan *formatter.Formatter

	responses = make(chan *types.TelegramResponse)
	requests = make(chan *formatter.Formatter)
	go pollResponse(responses)
	go dispatcher(responses, requests)
	go pushRequest(requests)

	for {
		time.Sleep(time.Second)
	}
}

func StartTests() {
	var responses chan *types.TelegramResponse
	var requests chan *formatter.Formatter

	responses = make(chan *types.TelegramResponse)
	requests = make(chan *formatter.Formatter)
	responses <- &types.TelegramResponse{
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

	go dispatcher(responses, requests)
	close(responses)
	r := <-requests
	r.AssertString("Hello! It's just a keyboard for a test, Jonh", true)
	r.AssertChatId(456, true)
	r.AssertInlineKeyboard([]int{1, 1}, []string{"I will send you a photo", "I will send you a video"}, []string{"/photo", "/video"}, []string{"cmd", "cmd"}, true)
	r.AssertPhoto("AgACAgQAAxkDAAIJRGW3rwaLqri1BkTdVQm1VFA8tE4HAAJeszEbEAABvFHW3MOANm9QFQEAAwIAA20AAzQE", true)
	r.AssertVideo("BAACAgIAAxkDAAIJW2W3sTguaruPGvo722qeKTcOPwvxAAIzPQACy-DASekiOEg76qGiNAQ", true)
	fmt.Print("All was alright!")
}
