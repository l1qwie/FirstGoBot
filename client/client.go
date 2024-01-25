package client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"strings"
)

func comStart(name string) (answer string) {
	return fmt.Sprint("Hello, World! Hello, ", name)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func CreateFormFile(fieldname, filename string, w *multipart.Writer) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	h.Set("Content-Type", "image/jpeg")
	return w.CreatePart(h)
}

func comPhoto(buf *bytes.Buffer, user_id int, contenttype *string) (err error) {
	var (
		file   *os.File
		part   io.Writer
		writer *multipart.Writer
	)

	writer = multipart.NewWriter(buf)
	file, err = os.Open("FOOTBALL1.jpg")
	if err == nil {
		part, err = writer.CreateFormFile("photo", "FOOTBALL1.jpg")
	}
	if err == nil {
		_, err = io.Copy(part, file)
	}
	if err == nil {
		err = writer.WriteField("chat_id", fmt.Sprintf("%d", user_id))
	}
	if err == nil {
		err = writer.Close()
	}
	*contenttype = writer.FormDataContentType()
	file.Close()
	return err
}

//func comKeyboard() {

//}

func handlerErrors(name string) (answer string) {
	return fmt.Sprint("Sorry, I couldn't understand you, ", name)
}

func Directioner(com, name string, user_id int, typeofmes, answer *string, buf *bytes.Buffer, contenttype *string) {
	if com == "/start" {
		*typeofmes = "mesText"
		*answer = comStart(name)
	} else if com == "/photo" {
		*typeofmes = "mesPhoto"
		comPhoto(buf, user_id, contenttype)
		//} else if com == "/keyboard" {
		//	*typeofmes = "mesKeyboard"
		//	comKeyboard()
	} else {
		*typeofmes = "mesText"
		*answer = handlerErrors(name)
	}
}
