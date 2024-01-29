package executer

import (
	"encoding/json"
	types "firstgobot/byogram/types"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Updates(TelebotToken string, offset, user_id *int, text, name *string) (err error) {
	var (
		url      string
		response *http.Response
		body     []byte
	)

	url = fmt.Sprintf(types.HttpsRequest+"bot%s/getUpdates?limit=1&offset=%d", TelebotToken, *offset)
	response, err = http.Get(url)
	if err == nil {
		defer response.Body.Close()
		body, err = io.ReadAll(response.Body)
		if err == nil {
			err = handleTelegramResponse(body, text, name, user_id)
		}
	}

	return err
}

func handleTelegramResponse(response []byte, text, name *string, user_id *int) (err error) {
	var telegramAnswer types.TelegramAnswer
	err = json.Unmarshal(response, &telegramAnswer)
	fmt.Println(fmt.Sprintf("Json-ответ от телеграма: %s", response))
	if err == nil {
		if !telegramAnswer.Ok {
			err = fmt.Errorf("Telegram API вернул ошибку: %s", telegramAnswer.Result)
		} else {
			if len(telegramAnswer.Result) > 0 {
				*text = telegramAnswer.Result[0].Message.Text
				*name = telegramAnswer.Result[0].Message.TypeFrom.Name
				*user_id = telegramAnswer.Result[0].Message.TypeFrom.UserID
				fmt.Println(fmt.Sprintf("Обработка успешного ответа от Telegram. Name: %s Text: %s UserID: %d", string(*name), string(*text), int(*user_id)))
			} else {
				err = fmt.Errorf("telegramAnswer.Result ничему не равен!")
			}
		}
	}

	return err
}

func FirstStep(response []byte, offset *int) (err error) {
	var telegramResponse types.TelegramResponse

	err = json.Unmarshal(response, &telegramResponse)
	if err == nil {
		// Проверка, что Result - это массив
		if len(telegramResponse.Result) > 0 {
			// Перебираем элементы массива
			for _, storage := range telegramResponse.Result {
				*offset = storage.ID
				fmt.Println("Обработка успешного ответа от Telegram. Update_id =", storage.ID)
			}
		} else {
			err = fmt.Errorf("Telegram API вернул пустой массив результатов")
		}
	}

	return err
}

func RequestOffset(TelebotToken string, offset *int) (returnederr error) {
	var (
		response *http.Response
		body     []byte
		err      error
	)

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
