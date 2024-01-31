package executer

import types "firstgobot/byogram/types"

func DoGetUpdates(token string, offset *int, telegramResponse *types.TelegramResponse) error {
	return Updates(token, offset, telegramResponse)
}

func HowToKnowOffset(token string, offset *int) error {
	return RequestOffset(token, offset)
}
