package main

import (
	"firstgobot/bot"
	executer "firstgobot/byogram/executer"
	"firstgobot/byogram/types"
	"fmt"
)

func main() {
	var (
		offset  int
		text    string
		name    string
		user_id int
		err     error
	)

	err = executer.HowToKnowOffset(types.TelebotToken, &offset)
	for err != nil {
		err = executer.HowToKnowOffset(types.TelebotToken, &offset)
	}
	for {
		err = executer.DoGetUpdates(types.TelebotToken, &offset, &user_id, &text, &name)
		if err != nil {
			fmt.Println("Ошибка при получении обновлений от Telegram:", err)
		} else {
			bot.Acceptance(text, name, user_id)
			//if err != nil {
			//	fmt.Println("Ошибка при попытке отправить сообщенрие пользователю:", err)
			//}
			offset = offset + 1
		}
	}
}

/*
	The main structur is:

	1. Executer (do) getUpdates

	2. Processing json response from Telegram

	3. Executer redirections decrypted data from json to Client

	4. Client do something what client wants to do

	5. (Optional) Client can do something with database

	6. Client returns to Executer what action executer should will do

	7. Executer (do) HTTPS-request
*/
