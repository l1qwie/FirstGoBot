package executer

import (
	client "firstgobot/client"
	"fmt"
	"net/http"
)

func sendMessage(TelebotToken string, text string, name *string) (returnederr error) {
	var answer, url string
	var err error
	var response *http.Response

	answer = client.ComStart(text, name)

	url = fmt.Sprintf(httpsRequest+"bot%s/sendMessage?chat_id=%d&text=%s", TelebotToken, 738070596, answer)

	response, err = http.Get(url)
	if err != nil {
		returnederr = err
	} else {
		defer response.Body.Close()
	}
	return returnederr
}
