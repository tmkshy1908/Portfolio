package interfaces

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
	"github.com/tmkshy1908/Portfolio/usecase"
)

type CommonController struct {
	controller Controller
	Interactor CommonInteractor
}

type Controller interface {
	Sayhello(http.ResponseWriter, *http.Request)
	SampleHandler(http.ResponseWriter, *http.Request)
}

func NewController(SqlHandler db.SqlHandler) (cc *CommonController) {
	// UseCase interface 構造体の値の初期化
	cc = &CommonController{
		Interactor: &usecase.CommonInteractor{
			CommonRepository: &CommonRepository{
				DB: SqlHandler,
			},
		},
	}
	return
}

func (cc *CommonController) Sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	fmt.Println("SayhelloName")
}

func (cc *CommonController) SampleHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()

	resp, err := cc.Interactor.UseCaseSampleRepository(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
