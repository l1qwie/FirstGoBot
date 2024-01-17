package main

import (
	Executer "firstgobot/Executer"
	"fmt"
)

func main() {
	var offset int
	var text string
	var name string
	var err error

	err = Executer.HowToKnowOffset(TelebotToken, &offset)
	if err != nil {
		fmt.Println("Возникла ошибка при попытке узнать offset:", err)
	} else {
		for true {
			err = Executer.DoGetUpdates(TelebotToken, &offset, &text, &name)
			if err != nil {
				fmt.Println("Ошибка при получении обновлений от Telegram:", err)
			} else {
				err = Executer.Redirectioner(TelebotToken, text, &name)
				if err != nil {
					fmt.Println("Ошибка при попытке отправить сообщенрие пользователю:", err)
				}
			}
			offset = offset + 1
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
}
