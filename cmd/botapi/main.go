package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/line"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	bot, err := line.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	db, err := db.NewHandler()
	if err != nil {
		fmt.Println(err)
	}
	infrastructure.NewServer(db, bot)
}
