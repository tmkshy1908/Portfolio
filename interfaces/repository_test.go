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

var i usecase.CommonRepository
var contents *domain.Contents
var bot line.Client
var Db db.SqlHandler
var err error
var userId string

func Newinit() {
	// var err error
	err = godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}
	bot, err = line.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	Db, err = db.NewHandler()
	if err != nil {
		fmt.Println(err)
	}
	i = &interfaces.CommonRepository{DB: Db, Bot: bot}
	userId = "U024602d41dff0396b1630a44ccd5dbee"

	// contents = &domain.Contents{Contents_Day: tt, Location: "池袋", EventTitle: "イベントイベント", Act: "吉田ヨシダ　田中たなか　岡田okada", OtherInfo: "19:00 end"}
}

func Test_Add(t *testing.T) {
	Newinit()
	ctx := context.Background()
	str := "2020/08/20"
	layout := "2006/01/02"
	tt, _ := time.Parse(layout, str)
	contents = &domain.Contents{Contents_Day: tt, Location: "池袋", EventTitle: "イベントイベント", Act: "吉田ヨシダ　田中たなか　岡田okada", OtherInfo: "19:00 end"}
	err = i.Add(ctx, contents)
	if err != nil {
		fmt.Println(err)
	}
}

func Test_Find(t *testing.T) {
	Newinit()
	ctx := context.Background()
	a, err := i.Find(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func Test_Update(t *testing.T) {
	Newinit()
	ctx := context.Background()
	str := "2020/08/20"
	layout := "2006/01/02"
	tt, _ := time.Parse(layout, str)
	contents = &domain.Contents{Contents_Day: tt, Location: "池袋", EventTitle: "アップデート", Act: "アプデーと", OtherInfo: "19:00 end"}
	i.Update(ctx, contents)
}

func Test_Delete(t *testing.T) {
	Newinit()
	ctx := context.Background()
	str := "2020/08/20"
	layout := "2006/01/02"
	tt, _ := time.Parse(layout, str)
	contents = &domain.Contents{Contents_Day: tt, Location: "池袋", EventTitle: "イベントイベント", Act: "吉田ヨシダ　田中たなか　岡田okada", OtherInfo: "19:00 end"}
	i.Delete(ctx, contents)
}

func Test_CallReply(t *testing.T) {
	Newinit()
	msg := "テストテスト"

	i.CallReply(msg, userId)

}

func Test_WaitMsg(t *testing.T) {
	Newinit()
	ctx := context.Background()
	contents, err := i.WaitMsg(ctx)
	fmt.Println(contents, err)
}

func TestStartUser(t *testing.T) {
	Newinit()
	ctx := context.Background()
	i.StartUser(ctx, userId)
}

func Test_UserCheck(t *testing.T) {
	Newinit()
	ctx := context.Background()
	if i.UserCheck(ctx, userId) {
		fmt.Println("登録ずみ")
	} else {
		fmt.Println("未登録")
	}
}

func Test_EndUser(t *testing.T) {
	Newinit()
	ctx := context.Background()
	i.EndUser(ctx, userId)
}
