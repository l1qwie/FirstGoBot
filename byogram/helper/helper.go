package helper

import "firstgobot/byogram/types"

func ReturnText(telegramResponse types.TelegramResponse) (text string) {
	if telegramResponse.Result[0].Message.Text != "" {
		text = telegramResponse.Result[0].Message.Text
	} else if telegramResponse.Result[0].Query.Data != "" {
		text = telegramResponse.Result[0].Query.Data
	}
	return text
}

func ReturnChatId(telegramResponse types.TelegramResponse) (chatID int) {
	if telegramResponse.Result[0].Message.TypeFrom.UserID != 0 {
		chatID = telegramResponse.Result[0].Message.TypeFrom.UserID
	} else if telegramResponse.Result[0].Query.TypeFrom.UserID != 0 {
		chatID = telegramResponse.Result[0].Query.TypeFrom.UserID
	}
	return chatID
}

func ReturnName(telegramResponse types.TelegramResponse) (name string) {
	if telegramResponse.Result[0].Message.TypeFrom.Name != "" {
		name = telegramResponse.Result[0].Message.TypeFrom.Name
	} else if telegramResponse.Result[0].Query.TypeFrom.Name != "" {
		name = telegramResponse.Result[0].Query.TypeFrom.Name
	}
	return name
}

func ReturnLastName(telegramResponse types.TelegramResponse) (lastname string) {
	if telegramResponse.Result[0].Message.TypeFrom.LastName != "" {
		lastname = telegramResponse.Result[0].Message.TypeFrom.LastName
	} else if telegramResponse.Result[0].Query.TypeFrom.LastName != "" {
		lastname = telegramResponse.Result[0].Query.TypeFrom.LastName
	}
	return lastname
}

func ReturnUsername(telegramResponse types.TelegramResponse) (username string) {
	if telegramResponse.Result[0].Message.TypeFrom.Username != "" {
		username = telegramResponse.Result[0].Message.TypeFrom.Username
	} else if telegramResponse.Result[0].Query.TypeFrom.Username != "" {
		username = telegramResponse.Result[0].Query.TypeFrom.Username
	}
	return username
}

func ReturnBotStatus(telegramResponse types.TelegramResponse) (botstatus bool) {

	m_isbot := telegramResponse.Result[0].Message.TypeFrom.IsBot
	cl_isbot := telegramResponse.Result[0].Query.TypeFrom.IsBot

	if !m_isbot && !cl_isbot {
		botstatus = false
	} else if m_isbot && !cl_isbot || !m_isbot && cl_isbot {
		botstatus = true
	}
	return botstatus
}
