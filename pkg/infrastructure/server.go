package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/tmkshy1908/Portfolio/interfaces"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/line"
)

type ControllHandler struct {
	CommonController *interfaces.CommonController
	// 実態にアクセスするために*を使う
}

func NewServer(h db.SqlHandler, b line.LineClient) {
	c := &ControllHandler{
		CommonController: interfaces.NewController(h, b),
	}

	NewRouter(c)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("serverOK")
	}
}
