package database

import (
	"database/sql"
	"firstgobot/byogram/types"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Game struct {
	ID    int
	Sport string
	Date  int
	Time  int
}

func FirstConnect() (err error) {
	var (
		db    *sql.DB
		rows  *sql.Rows
		games []Game
		game  Game
	)

	db, err = sql.Open("postgres", types.ConnectTo())
	if err == nil {
		fmt.Println("Connected to database")
		rows, err = db.Query("SELECT game_id, sport, date, time FROM Schedule")
	}
	if err == nil {
		for rows.Next() {
			err = rows.Scan(&game.ID, &game.Sport, &game.Date, &game.Time)
			if err != nil {
				log.Fatal(err)
			}
			games = append(games, game)
		}
	}
	fmt.Println(games)
	if err == nil {
		err = rows.Err()
	}

	fmt.Println("Список игр:")
	for _, game := range games {
		fmt.Printf("ID: %d, Sport: %s, Date: %d, Time: %d\n", game.ID, game.Sport, game.Date, game.Time)
	}
	db.Close()

	return err
}
