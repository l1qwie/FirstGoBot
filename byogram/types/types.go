package types

const HttpsRequest = "https://api.telegram.org/"

type TelegramResponse struct {
	Ok     bool            `json:"ok"`
	Result []StorageOfJson `json:"result"`
}

type InlineKeyboardMarkup struct {
	Kb []InlineKeyboardButton
}

type InlineKeyboardButton struct {
	Text          string
	Url           string
	Callback_data string
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
