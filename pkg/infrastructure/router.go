package infrastructure

import (
	"net/http"
)

func NewRouter(controller *ControllHandler) {

	http.HandleFunc("/", controller.CommonController.SayhelloName)
}

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello")

// }
