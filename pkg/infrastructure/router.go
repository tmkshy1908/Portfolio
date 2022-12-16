package infrastructure

import (
	"net/http"
)

func NewRouter(controller *ControllHandler) {
	http.HandleFunc("/hello", controller.CommonController.Sayhello)
	http.HandleFunc("/sample", controller.CommonController.SampleHandler)
	http.HandleFunc("/line", controller.CommonController.LineHandller)
}
