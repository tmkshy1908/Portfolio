package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/tmkshy1908/Portfolio/interfaces"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
)

type ControllHandler struct {
	CommonController *interfaces.CommonController
	// 実態にアクセスするために*を使う
}

func NewServer(h db.SqlHandler) {
	c := &ControllHandler{
		CommonController: interfaces.NewController(h),
	}

	NewRouter(c)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
