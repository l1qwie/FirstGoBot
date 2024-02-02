package main

import (
	"firstgobot/database"
	"log"
)

func main() {
	err := database.FirstConnect()
	if err != nil {
		log.Fatal(err)
	}
	//byogram.StartWithTelegram()
	//byogram.StartTests()
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
