package main

import (
	"fmt"

	"github.com/tmkshy1908/Portfolio/pkg/infrastructure"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
)

func main() {
	db, err := db.NewHandler()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(db)
	}

	// srv := &http.Server{
	// 	Addr: ,
	// 	Handler: ,
	// }

	infrastructure.NewServer()

}
