package executer

import (
	"encoding/json"
	types "firstgobot/byogram/types"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Updates(TelebotToken string, offset *int, telegramResponse *types.TelegramResponse) (err error) {
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
			err = handleTelegramResponse(body, telegramResponse)
		}
	}
	/*if len(telegramResponse.Result) == 0 {
		fmt.Println(url)
		fmt.Println(string(body))
		log.Fatal(response)
	}*/
	return err
}

func handleTelegramResponse(response []byte, telegramResponse *types.TelegramResponse) (err error) {
	err = json.Unmarshal(response, &telegramResponse)
	if err == nil {
		if !telegramResponse.Ok {
			err = fmt.Errorf(fmt.Sprintf("Telegram API вернул ошибку: %s", telegramResponse.Error.Message))
		}
	}

	return err
}

func FirstStep(response []byte, offset *int) (err error) {
	var telegramResponse types.JustForUpdate

	err = json.Unmarshal(response, &telegramResponse)
	if err == nil {
		// Проверка, что Result - это массив
		if len(telegramResponse.Result) > 0 {
			// Перебираем элементы массива
			for _, storage := range telegramResponse.Result {
				*offset = storage.ID
				//fmt.Println("Обработка успешного ответа от Telegram. Update_id =", storage.ID)
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
