package byogram

import (
	"firstgobot/bot"
	"firstgobot/byogram/executer"
	"firstgobot/byogram/formatter"
	"firstgobot/byogram/helper"
	"firstgobot/byogram/types"
	"fmt"
	"log"
	"time"
)

func pollResponse(output chan<- *formatter.Formatter, reg *executer.RegTable) {
	var (
		offset           int
		err              error
		telegramResponse *types.TelegramResponse
		index            int
		chatID           int
		//privateChan      chan *types.TelegramResponse
	)

	err = executer.RequestOffset(types.TelebotToken, &offset)
	for err != nil {
		err = executer.RequestOffset(types.TelebotToken, &offset)
	}
	//privateChan = make(chan *types.TelegramResponse)
	for {
		telegramResponse = new(types.TelegramResponse)
		err = executer.Updates(&offset, telegramResponse)
		if len(telegramResponse.Result) != 0 && err == nil {
			chatID = helper.ReturnChatId(telegramResponse)

			index = reg.Seeker(chatID)
			if index != executer.None {
				reg.Reg[index].Ch <- telegramResponse
			} else {
				index = reg.NewIndex()
				reg.Reg[index].UserId = chatID
				reg.Reg[index].Ch = make(chan *types.TelegramResponse, 10)
				fmt.Println("help1", index)
				reg.Reg[index].Ch <- telegramResponse
				fmt.Println("help2")
				go worker(reg.Reg[index].Ch, output)
			}
			offset = offset + 1
		} else if err != nil {
			log.Fatal(err)
		}
	}
}

func worker(input <-chan *types.TelegramResponse, output chan<- *formatter.Formatter) {
	fmt.Println("In worker")
	for len(input) > 0 {
		//for r, any := <- input {
		fm := bot.Receiving(<-input)
		fm.Complete()
		output <- fm
	}
	//закрой за мной дверь, я ухожу
}

func pushRequest(requests <-chan *formatter.Formatter) {
	for r := range requests {
		err := r.Send()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func StartWithTelegram() {
	var (
		requests chan *formatter.Formatter
		reg      *executer.RegTable
	)
	requests = make(chan *formatter.Formatter)
	reg = new(executer.RegTable)
	go pollResponse(requests, reg)
	//go dispatcher(responses, requests)
	go pushRequest(requests)
	//go doubleСheck(responses, stg)

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

	//go dispatcher(responses, requests)
	close(responses)
	r := <-requests
	r.AssertString("Hello! It's just a keyboard for a test, Jonh", true)
	r.AssertChatId(456, true)
	r.AssertInlineKeyboard([]int{1, 1}, []string{"I will send you a photo", "I will send you a video"}, []string{"/photo", "/video"}, []string{"cmd", "cmd"}, true)
	r.AssertPhoto("AgACAgQAAxkDAAIJRGW3rwaLqri1BkTdVQm1VFA8tE4HAAJeszEbEAABvFHW3MOANm9QFQEAAwIAA20AAzQE", true)
	r.AssertVideo("BAACAgIAAxkDAAIJW2W3sTguaruPGvo722qeKTcOPwvxAAIzPQACy-DASekiOEg76qGiNAQ", true)
	fmt.Print("All was alright!")
}
