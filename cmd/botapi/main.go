package main

import (
	"fmt"
	"net/http"

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
	http.HandleFunc("/", sayhelloName)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}

}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
