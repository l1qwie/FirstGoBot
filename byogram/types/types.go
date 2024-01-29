package types

import (
	"encoding/json"
	"log"
)

const HttpsRequest = "https://api.telegram.org/"

type SendMessagePayload struct {
	ChatID      int    `json:"chat_id"`
	Text        string `json:"text"`
	ReplyMarkup string `json:"reply_markup"`
	Photo       string `json:"photo"`
}

type MessagePhoto struct {
	ChatID  int    `json:"chat_id"`
	Caption string `json:"caption"`
	Photo   string `json:"photo"`
}

type TelegramResponse struct {
	Ok     bool            `json:"ok"`
	Result []StorageOfJson `json:"result"`
}

type InlineKeyboardMarkup struct {
	Buttons [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type StorageOfJson struct {
	ID int `json:"update_id"`
}

type TelegramAnswer struct {
	Ok     bool             `json:"ok"`
	Result []TelegramUpdate `json:"result"`
}

type TelegramUpdate struct {
	UpdateID int        `json:"update_id"`
	Message  InfMessage `json:"message"`
}

type InfMessage struct {
	TypeFrom FromUser `json:"from"`
	Text     string   `json:"text"`
}

type FromUser struct {
	UserID int    `json:"id"`
	Name   string `json:"first_name"`
}

func CreateInlineKeyoard(keyboard InlineKeyboardMarkup) string {
	var (
		jsonData []byte
		err      error
	)

	jsonData, err = json.Marshal(keyboard)
	if err != nil {
		log.Fatal(err)
	}

	return string(jsonData)
}
