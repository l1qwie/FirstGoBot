package formatter

import (
	"bytes"
	"encoding/json"
	"firstgobot/byogram/methods"
	"firstgobot/byogram/types"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

type Formatter struct {
	//TypeOfMessage string
	Message types.SendMessagePayload
	//ImageMessage types.MessagePhoto
	Keyboard InlineKeyboard
	//photo        photo
	contenttype string
}

type InlineKeyboard struct {
	Keyboard [][]btn `json:"inline_keyboard"`
	//coordinates []int
	x int
	y int
}

type btnKind int

const (
	bCmd btnKind = 1
	bUrl btnKind = 2
)

type btn struct {
	Label string `json:"text"`
	what  btnKind
	Cmd   string `json:"callback_data"`
	Url   string `json:"url"`
}

type photo struct {
	buf         bytes.Buffer
	contenttype string
}

func (fm *Formatter) WriteString(lineoftext string) {
	fm.Message.Text = lineoftext
}

func (fm *Formatter) WriteChatId(chatID int) {
	fm.Message.ChatID = chatID
}

func (fm *Formatter) SetIkbdDim(dim []int) {

	fm.Keyboard.Keyboard = make([][]btn, len(dim))
	for i := 0; i < len(dim); i++ {
		fm.Keyboard.Keyboard[i] = make([]btn, dim[i])
	}
	//fm.Keyboard.coordinates = dim
}

func (fm *Formatter) doRutine() {
	//if fm.Keyboard.x == fm.Keyboard.coordinates[fm.Keyboard.y] {
	if fm.Keyboard.x == len(fm.Keyboard.Keyboard[fm.Keyboard.y]) {
		fm.Keyboard.x = 0
		fm.Keyboard.y = fm.Keyboard.y + 1
		//i := len(fm.Keyboard.Keyboard)
		//len(fm.Keyboard.Keyboard[i])
	}
}

func (fm *Formatter) WriteInlineButtonCmd(label, cmd string) {
	fm.doRutine()
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Label = label
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].what = bCmd
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Cmd = cmd

	fm.Keyboard.x = fm.Keyboard.x + 1

}

func (fm *Formatter) WriteInlineButtonUrl(label, url string) {
	fm.doRutine()
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Label = label
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].what = bUrl
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Url = url

	fm.Keyboard.x = fm.Keyboard.x + 1

}

/*
func (fm *Formatter) crossСheck() {
	if fm.ImageMessage.Caption == "" {
		fm.ImageMessage.Caption = fm.Message.Text
	}
	if fm.ImageMessage.ChatID == 0 {
		fm.ImageMessage.ChatID = fm.Message.ChatID
	}
}
*/

func (fm *Formatter) AddPhotoFromMemmory(path string) {
	fm.Message.Photo = path
}

func (fm *Formatter) AddPhotoFromInternet(path string) {

}

func (fm *Formatter) AddPhotoFromTG(path string) {

}

func (fm *Formatter) PreparePhoto(buf *bytes.Buffer) (string, error) {
	var (
		file   *os.File
		part   io.Writer
		writer *multipart.Writer
		err    error
	)

	writer = multipart.NewWriter(buf)
	file, err = os.Open(fm.Message.Photo)
	if err == nil {
		part, err = writer.CreateFormFile("photo", fm.Message.Photo)
	}
	if err == nil {
		_, err = io.Copy(part, file)
	}
	if err == nil {
		err = writer.WriteField("chat_id", fmt.Sprintf("%d", fm.Message.ChatID))
	}
	if err == nil {
		err = writer.WriteField("caption", fm.Message.Text)
	}
	if err == nil {
		err = writer.Close()
	}
	file.Close()

	return writer.FormDataContentType(), err

}

func (fm *Formatter) SendMessage() {
	var (
		jsonMessage  []byte
		jsonKeyboard []byte
		err          error
		finalBuffer  *bytes.Buffer
		function     string
	)

	if fm.Keyboard.Keyboard != nil {
		jsonKeyboard, err = json.Marshal(fm.Keyboard)
		if err != nil {
			log.Fatal(err)
		}
		fm.Message.ReplyMarkup = string(jsonKeyboard)
	}

	if fm.Message.Photo == "" {
		jsonMessage, err = json.Marshal(fm.Message)
		function = "sendMessage"
		fm.contenttype = "application/json"
		finalBuffer = bytes.NewBuffer(jsonMessage)
	} else {
		function = "sendPhoto"
		finalBuffer = bytes.NewBuffer(nil)
		fm.contenttype, err = fm.PreparePhoto(finalBuffer)
		if err != nil {
			log.Fatal(err)
		}
	}
	//fmt.Println(fmt.Println(finalBuffer.String()))
	if err != nil {
		log.Fatal(err)
	}

	methods.Send(finalBuffer, function, fm.contenttype)
}
