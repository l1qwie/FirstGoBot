package types

const HttpsRequest = "https://api.telegram.org/"

/*
 */
type InfMessage struct {
	TypeFrom User   `json:"from"`
	Text     string `json:"text"`
}

type User struct {
	UserID   int    `json:"id"`
	IsBot    bool   `json:"is_bot"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
}

type TelegramUpdate struct {
	UpdateID int        `json:"update_id"`
	Message  InfMessage `json:"message"`
	Query    Callback   `json:"callback_query"`
}

type Callback struct {
	TypeFrom User   `json:"from"`
	Data     string `json:"data"`
}

type TelegramError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TelegramResponse struct {
	Ok     bool             `json:"ok"`
	Result []TelegramUpdate `json:"result"`
	Error  *TelegramError   `json:"error,omitempty"`
}

type JustForUpdate struct {
	Ok     bool            `json:"ok"`
	Result []StorageOfJson `json:"result"`
}

type StorageOfJson struct {
	ID int `json:"update_id"`
}

type SendMessagePayload struct {
	ChatID      int    `json:"chat_id"`
	Text        string `json:"text"`
	ReplyMarkup string `json:"reply_markup"`
	Photo       string `json:"photo"`
	Video       string `json:"video"`
}

type FMTRS interface {
	WriteString(string)
	WriteChatId(int)
	AddPhotoFromStorage(string)
	AddPhotoFromTG(string)
	AddPhotoFromInternet(string)
	AddVideoFromStorage(string)
	AddVideoFromTG(string)
	AddVideoFromInternet(string)
	SetIkbdDim([]int)
	WriteInlineButtonCmd(string, string)
	WriteInlineButtonUrl(string, string)
}
