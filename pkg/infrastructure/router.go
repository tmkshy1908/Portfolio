package infrastructure

import (
	"fmt"
	"net/http"
)

func NewRouter(controller *ControllHandler) {
	fmt.Println("NewRouter")
	http.HandleFunc("/hello", controller.CommonController.Sayhello)
	http.HandleFunc("/sample", controller.CommonController.SampleHandler)
}
