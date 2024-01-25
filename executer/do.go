package executer

import (
	"bytes"
	"encoding/json"
	client "firstgobot/client"
	"fmt"
	"log"
	"net/http"
)

func getGet(url string) (err error) {
	var response *http.Response
	response, err = http.Get(url)
	if err == nil {
		defer response.Body.Close()
	}

	return err
}

func formatJSON(data []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", " ")

	if err != nil {
		fmt.Println(err)
	}

	d := out.Bytes()
	return string(d)
}

func getPost(url string, user_id int, Buffer *bytes.Buffer, contenttype string) (err error) {
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
	log.Fatal(err)

	return err
}

func sendMessage(TelebotToken, text, name string, user_id int) (err error) {
	var (
		answer, url string
		Buffer      bytes.Buffer
		typeofmes   string
		contenttype string
	)

	client.Directioner(text, name, user_id, &typeofmes, &answer, &Buffer, &contenttype)

	if typeofmes == "mesText" {
		url = fmt.Sprintf(httpsRequest+"bot%s/sendMessage?chat_id=%d&text=%s", TelebotToken, user_id, answer)
		err = getGet(url)
	} else if typeofmes == "mesPhoto" {
		url = fmt.Sprintf(httpsRequest+"bot%s/sendPhoto", TelebotToken)
		err = getPost(url, user_id, &Buffer, contenttype)
	}

	return err
}

//738070596
