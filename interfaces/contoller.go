package interfaces

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/line"
	"github.com/tmkshy1908/Portfolio/usecase"
)

type CommonController struct {
	// controller Controller
	Interactor CommonInteractor
}

type Controller interface {
	Sayhello(http.ResponseWriter, *http.Request)
	SampleHandler(http.ResponseWriter, *http.Request)
	LineHandller(http.ResponseWriter, *http.Request)
}

func NewController(SqlHandler db.SqlHandler, LineHandller line.LineClient) (cc *CommonController) {
	// UseCase interface 構造体の値の初期化
	cc = &CommonController{
		Interactor: &usecase.CommonInteractor{
			CommonRepository: &CommonRepository{
				DB:  SqlHandler,
				Bot: LineHandller,
			},
		},
	}
	return
}

func (cc *CommonController) Sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func (cc *CommonController) LineHandller(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "request", r)
	cc.Interactor.DivideMessage(ctx)
}
