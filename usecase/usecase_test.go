package usecase_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/tmkshy1908/Portfolio/interfaces"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/line"
	"github.com/tmkshy1908/Portfolio/usecase"
)

type test_st struct {
	CommonInteractor interfaces.CommonInteractor
}

func Newinit() (ts *test_st) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}
	bot, err := line.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	db, err := db.NewHandler()
	if err != nil {
		fmt.Println(err)
	}

	ts = &test_st{
		CommonInteractor: &usecase.CommonInteractor{
			CommonRepository: &interfaces.CommonRepository{DB: db, Bot: bot},
		},
	}
	return
}

func Test_DivideMessage(t *testing.T) {
	// r := &http.Request{}

	reqBody := bytes.NewBufferString("request body")
	req := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)
	ts := Newinit()
	ctx := context.Background()
	ts.CommonInteractor.DivideMessage(ctx, req)
}
