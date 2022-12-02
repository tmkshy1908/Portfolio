package interfaces

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/tmkshy1908/Portfolio/usecase"
)

type CommonController struct {
	Interactor       Controller
	CommonInteractor usecase.CommonInteractor
	Converter        ConvertController
}

type Controller interface {
	Sayhello(http.ResponseWriter, *http.Request)
}

func NewController() Controller {
	return &CommonController{}
}

func (cc *CommonController) Sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	fmt.Println("SayhelloName")
}

func (cc *CommonController) SampleHandler(w http.ResponseWriter, r *http.Request) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()

	resp, err := cc.CommonInteractor.UseCaseSampleRepository(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	cc.Converter.ToSampleResponseData(resp)

	return
}
