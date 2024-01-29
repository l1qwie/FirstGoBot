package formatter

import "firstgobot/byogram/types"

type Formatter struct {
	Message     types.SendMessagePayload
	Keyboard    InlineKeyboard
	contenttype string
	kindofmedia int
	mediatype   string
}

type InlineKeyboard struct {
	Keyboard [][]btn `json:"inline_keyboard"`
	x        int
	y        int
}

type btnKind int

const (
	bCmd btnKind = 1
	bUrl btnKind = 2

	fromStorage  = 0
	fromTelegram = 1
	fromInternet = 2
)

type btn struct {
	Label string `json:"text"`
	what  btnKind
	Cmd   string `json:"callback_data"`
	Url   string `json:"url"`
}
