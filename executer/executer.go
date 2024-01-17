package executer

/*import (
	"firstgobot/executer"
)*/

func DoGetUpdates(token string, offset *int, text *string, name *string) error {
	return Updates(token, offset, text, name)
}

func HowToKnowOffset(token string, offset *int) error {
	return RequestOffset(token, offset)
}

func Redirectioner(token string, messagetext string, name *string) error {
	return sendMessage(token, messagetext, name)
}
