package executer

func DoGetUpdates(token string, offset, user_id *int, text, name *string) error {
	return Updates(token, offset, user_id, text, name)
}

func HowToKnowOffset(token string, offset *int) error {
	return RequestOffset(token, offset)
}

func Redirectioner(token, messagetext, name string, user_id int) error {
	return sendMessage(token, messagetext, name, user_id)
}
