package methods

import (
	"bytes"
	"firstgobot/byogram/executer"
	"firstgobot/byogram/types"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func imageRutine(buf *bytes.Buffer, image string, chatID int) (err error, contenttype string) {
	var (
		file   *os.File
		part   io.Writer
		writer *multipart.Writer
	)

	writer = multipart.NewWriter(buf)
	file, err = os.Open(image)
	if err == nil {
		part, err = writer.CreateFormFile("photo", image)
	}
	if err == nil {
		_, err = io.Copy(part, file)
	}
	if err == nil {
		err = writer.WriteField("chat_id", fmt.Sprintf("%d", chatID))
	}
	if err == nil {
		err = writer.Close()
	}
	contenttype = writer.FormDataContentType()
	file.Close()
	return err, contenttype
}

func SendMessage(text string, chatID int) {
	var (
		url string
		err error
	)

	url = fmt.Sprintf(types.HttpsRequest+"bot%s/sendMessage?chat_id=%d&text=%s", types.TelebotToken, chatID, text)
	err = executer.GetgetRequest(url)
	if err != nil {
		log.Fatal(err)
		//fmt.Println("Couldn't send the message. Error: ", err)
	}
}

func SendPhoto(image string, chatID int) {
	var (
		url         string
		err         error
		buf         bytes.Buffer
		contenttype string
	)

	url = fmt.Sprintf(types.HttpsRequest+"bot%s/sendPhoto", types.TelebotToken)
	err, contenttype = imageRutine(&buf, image, chatID)
	if err == nil {
		err = executer.GetpostRequest(url, &buf, contenttype)
		if err != nil {
			log.Fatal(err)
			//fmt.Println("Couldn't send the Image. Error: ", err)
		}
	} else {
		log.Fatal(err)
		//fmt.Println("Сouldn't prepare the Image for sending. Error: ", err)
	}

}
