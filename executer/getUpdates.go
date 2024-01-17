package executer

import (
	"encoding/json"
	types "firstgobot/types"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const httpsRequest = "https://api.telegram.org/"

func Updates(TelebotToken string, offset *int, text *string, name *string) (returnederr error) {
	var url string
	var response *http.Response
	var body []byte
	var err error

	url = fmt.Sprintf(httpsRequest+"bot%s/getUpdates?limit=1&offset=%d", TelebotToken, *offset)

	response, err = http.Get(url)
	if err != nil {
		returnederr = err
	} else {
		defer response.Body.Close()
		body, err = io.ReadAll(response.Body)
		if err != nil {
			returnederr = err
		} else {
			err = handleTelegramResponse(body, text, name)
			if err != nil {
				returnederr = err
			}
		}
	}

	return returnederr
}

func handleTelegramResponse(response []byte, text *string, name *string) (returnederr error) {
	var telegramAnswer types.TelegramAnswer
	var err error
	err = json.Unmarshal(response, &telegramAnswer)
	if err != nil {
		returnederr = err
	} else if !telegramAnswer.Ok {
		returnederr = fmt.Errorf("Telegram API вернул ошибку: %s", telegramAnswer.Result)
	} else if len(telegramAnswer.Result) > 0 {
		*text = telegramAnswer.Result[0].Message.Text
		*name = telegramAnswer.Result[0].Message.TypeFrom.Name

		str := fmt.Sprintf("Обработка успешного ответа от Telegram. Name: %s Text: %s", string(*name), string(*text))
		fmt.Println(str)
	} else {
		returnederr = fmt.Errorf("telegramAnswer.Result ничему не равен!")
	}
	return returnederr
}

func FirstStep(response []byte, offset *int) (returnederr error) {
	var telegramResponse types.TelegramResponse
	var err error

	err = json.Unmarshal(response, &telegramResponse)
	if err != nil {
		returnederr = err
	} else {
		// Проверка, что Result - это массив
		if len(telegramResponse.Result) > 0 {
			// Перебираем элементы массива
			for _, storage := range telegramResponse.Result {
				*offset = storage.ID
				fmt.Println("Обработка успешного ответа от Telegram. Update_id =", storage.ID)
			}
		} else {
			returnederr = fmt.Errorf("Telegram API вернул пустой массив результатов")
		}
	}

	return returnederr
}

func RequestOffset(TelebotToken string, offset *int) (returnederr error) {
	//var url string
	var response *http.Response
	var body []byte
	var err error

	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?limit=1", url.PathEscape(TelebotToken))
	response, err = http.Get(url)
	if err != nil {
		returnederr = err
	} else {
		defer response.Body.Close()
		body, err = io.ReadAll(response.Body)
		if err != nil {
			returnederr = err
		} else {
			err = FirstStep(body, offset)
			if err != nil {
				returnederr = err
			}
		}
	}

	return returnederr
}
