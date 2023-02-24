package interfaces_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tmkshy1908/Portfolio/domain"
	"github.com/tmkshy1908/Portfolio/interfaces"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/line"
	"github.com/tmkshy1908/Portfolio/usecase"
)

type test_st struct {
	commonRepository usecase.CommonRepository
	contents         *domain.Contents
	userId           string
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
	Db, err := db.NewHandler()
	if err != nil {
		fmt.Println(err)
	}

	str := "2020/08/20"
	layout := "2006/01/02"
	tt, _ := time.Parse(layout, str)

	ts = &test_st{commonRepository: &interfaces.CommonRepository{DB: Db, Bot: bot},
		contents: &domain.Contents{Contents_Day: tt, Location: "池袋", EventTitle: "イベントイベント", Act: "吉田ヨシダ　田中たなか　岡田okada", OtherInfo: "19:00 end"},
		userId:   "U024602d41dff0396b1630a44ccd5dbee",
	}
	return
}

func Test_Add(t *testing.T) {
	ts := Newinit()
	ctx := context.Background()
	err := ts.commonRepository.Add(ctx, ts.contents)
	if err != nil {
		fmt.Println(err)
	}
}

func Test_Find(t *testing.T) {
	ts := Newinit()
	ctx := context.Background()
	a, err := ts.commonRepository.Find(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func Test_Update(t *testing.T) {
	ts := Newinit()
	ctx := context.Background()
	ts.contents.Act = "アップデート"
	contents := ts.contents
	ts.commonRepository.Update(ctx, contents)
}

func Test_Delete(t *testing.T) {
	ts := Newinit()
	ctx := context.Background()
	ts.commonRepository.Delete(ctx, ts.contents)
}

func Test_CallReply(t *testing.T) {
	ts := Newinit()
	msg := "テストテスト"

	ts.commonRepository.CallReply(msg, ts.userId)

}

func Test_WaitMsg(t *testing.T) {
	ts := Newinit()
	ctx := context.Background()
	contents, err := ts.commonRepository.WaitMsg(ctx)
	fmt.Println(contents, err)
}

func TestStartUser(t *testing.T) {
	ts := Newinit()
	ctx := context.Background()
	ts.commonRepository.StartUser(ctx, ts.userId)
}

func Test_UserCheck(t *testing.T) {
	ts := Newinit()
	ctx := context.Background()
	if ts.commonRepository.UserCheck(ctx, ts.userId) {
		fmt.Println("登録ずみ")
	} else {
		fmt.Println("未登録")
	}
}

func Test_EndUser(t *testing.T) {
	ts := Newinit()
	ctx := context.Background()
	ts.commonRepository.EndUser(ctx, ts.userId)
}

// aaaa
