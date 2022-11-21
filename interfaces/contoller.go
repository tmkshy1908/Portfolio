package interfaces

import (
	"fmt"
	"net/http"
)

type CommonController struct {
	Interactor Controller
}

type Controller interface {
	SayhelloName(http.ResponseWriter, *http.Request)
}

func NewController() Controller {
	return &CommonController{}
}

func (cc *CommonController) SayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
