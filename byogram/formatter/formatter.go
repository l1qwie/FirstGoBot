package formatter

import (
	"bytes"
	"encoding/json"
	"firstgobot/byogram/methods"
)

func (fm *Formatter) WriteString(lineoftext string) {
	fm.Message.Text = lineoftext
}

func (fm *Formatter) WriteChatId(chatID int) {
	fm.Message.ChatID = chatID
}

func (fm *Formatter) SendMessage() error {
	var (
		jsonMessage  []byte
		jsonKeyboard []byte
		err          error
		finalBuffer  *bytes.Buffer
		function     string
	)

	if fm.Keyboard.Keyboard != nil {
		jsonKeyboard, err = json.Marshal(fm.Keyboard)
		if err == nil {
			fm.Message.ReplyMarkup = string(jsonKeyboard)
		}
	}

	if err == nil {
		if fm.Message.Photo == "" && fm.Message.Video == "" {
			function = "sendMessage"
			jsonMessage, err = json.Marshal(fm.Message)
			if err == nil {
				fm.contenttype = "application/json"
				finalBuffer = bytes.NewBuffer(jsonMessage)
			}
		} else if fm.Message.Video != "" || fm.Message.Photo != "" {
			if fm.mediatype == "photo" {
				function = "sendPhoto"
			} else if fm.mediatype == "video" {
				function = "sendVideo"
			}

			if fm.kindofmedia == fromStorage {
				finalBuffer = bytes.NewBuffer(nil)
				fm.contenttype, err = fm.PrepareMedia(finalBuffer)
			} else {
				jsonMessage, err = json.Marshal(fm.Message)
				if err == nil {
					fm.contenttype = "application/json"
					finalBuffer = bytes.NewBuffer(jsonMessage)
				}
			}
		}
	}
	//fmt.Print("There is a list", fmt.Sprint(finalBuffer.String()))
	if err == nil {
		methods.Send(finalBuffer, function, fm.contenttype)
	}

	return err
}
