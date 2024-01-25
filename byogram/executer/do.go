package executer

import (
	"bytes"
	"net/http"
)

func GetgetRequest(url string) (err error) {
	var response *http.Response
	response, err = http.Get(url)
	if err == nil {
		defer response.Body.Close()
	}

	return err
}

func GetpostRequest(url string, Buffer *bytes.Buffer, contenttype string) (err error) {
	var (
		request  *http.Request
		response *http.Response
		client   *http.Client
	)

	request, err = http.NewRequest("POST", url, Buffer)
	if err == nil {
		request.Header.Set("Content-Type", contenttype)
		client = &http.Client{}
		response, err = client.Do(request)
	}
	response.Body.Close()

	return err
}

/*
func createAnswer(TelebotToken, text, name string, chatID int) (err error) {
	var (
		answer, url, typeofmes, contenttype string
		Buffer                              bytes.Buffer
		keyboard                            types.InlineKeyboardMarkup
		dataload                            map[string]interface{}
	)

	dataload = map[string]interface{}{
		"chat_id":      chatID,
		"text":         answer,
		"reply_markup": keyboard,
	}

	url = fmt.Sprintf("%sbot%s/%s", httpsRequest, TelebotToken, method)

	keyboard = client.Directioner(text, name, user_id, &typeofmes, &answer, &Buffer, &contenttype)

	if typeofmes == "mesText" {
		url = fmt.Sprintf(httpsRequest+"bot%s/sendMessage?chat_id=%d&text=%s", TelebotToken, user_id, answer)
		err = getGet(url)
	} else if typeofmes == "mesPhoto" {
		url = fmt.Sprintf(httpsRequest+"bot%s/sendPhoto", TelebotToken)
		err = getPost(url, user_id, &Buffer, contenttype)
	} else if typeofmes == "mesKeyboard" {
		url = fmt.Sprintf("%sbot%s/sendMessage?chat_id=%d&text=%s&reply_markup=%s", httpsRequest, TelebotToken, user_id, answer, keyboard)
		err = getGet(url)
		log.Fatal(err)
	}

	return err
}
*/
