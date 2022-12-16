package main

import (
	"fmt"

	"github.com/tmkshy1908/Portfolio/pkg/infrastructure"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/api"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
)

func main() {
	bot, err := api.NewLineHandller()
	if err != nil {
		fmt.Println(err)
	}
	db, err := db.NewHandler()
	if err != nil {
		fmt.Println(err)
	}
	infrastructure.NewServer(db, bot)

}
