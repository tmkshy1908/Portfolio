package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/tmkshy1908/Portfolio/interfaces"
)

type ControllHandler struct {
	CommonController interfaces.Controller
}

func NewServer() {
	c := &ControllHandler{
		CommonController: interfaces.NewController(),
	}

	NewRouter(c)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
