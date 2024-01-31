package executer

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetgetRequest(url string) (err error) {
	var response *http.Response
	response, err = http.Get(url)
	if err == nil {
		defer response.Body.Close()
	}
	_, err = io.ReadAll(response.Body)
	//fmt.Println(string(body))

	return err
}

func GetpostRequest(url string, Buffer *bytes.Buffer, contenttype string) (err error) {
	var (
		request  *http.Request
		response *http.Response
		client   *http.Client
	)

	request, err = http.NewRequest("POST", url, Buffer)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", contenttype)
	client = &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	return err
}
