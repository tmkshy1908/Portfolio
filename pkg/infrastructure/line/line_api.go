package line

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

// type apiConnection struct {
// 	ChannelSecrets string
// 	AccessToken    string
// }

type LineConf struct {
	Bot *linebot.Client
}

type Client interface {
	CathEvents(context.Context, *http.Request) (string, string)
	MsgReply(string, string)
	WaitEvents(context.Context) (string, string, string, string, string)
	CathID(*http.Request) string
	TestFunc(context.Context, *http.Request) string
}

func NewClient() (lh Client, err error) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("ACCESS_TOKEN"),
	)
	if err != nil {
		fmt.Println("linebot.Newエラー", err)
	} else {
		fmt.Println("Api Connected.")
	}
	// a := bot.GetProfile()
	// fmt.Println(a)
	lh = &LineConf{Bot: bot}

	return
}

func (bot *LineConf) CathEvents(ctx context.Context, req *http.Request) (msg string, userId string) {
	events, err := bot.Bot.ParseRequest(req)
	if err != nil {
		fmt.Println("ParseReq", err)
	}
	for _, event := range events {

		if event.Type == linebot.EventTypeMessage {
			userId = event.Source.UserID

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msg = message.Text
				if _, err := bot.Bot.PushMessage(userId, linebot.NewTextMessage(msg)).Do(); err != nil {
					fmt.Println(err, "プッシュエラー")
				}

			case *linebot.StickerMessage:
				bot.MsgReply(stickerReply, userId)

			case *linebot.ImageMessage:
				bot.MsgReply(imageReply, userId)
			}
		} else {
			fmt.Println("EventTypeが違う")
		}
	}
	return
}

const (
	stickerReply string = "いいスタンプだね！"
	imageReply   string = "素敵な写真だね！"
)

func (bot *LineConf) MsgReply(msg string, userId string) {
	if _, err := bot.Bot.PushMessage(userId, linebot.NewTextMessage(msg)).Do(); err != nil {
		fmt.Println(err, "プッシュエラー")
	}
	// replyMessage := linebot.NewTextMessage(msg)
	// bot.Bot.BroadcastMessage(replyMessage).Do()
}

func (bot *LineConf) WaitEvents(ctx context.Context) (day string, location string, title string, act string, info string) {
	// c := &domain.Contents{}
	day = "22-04-09 00:00:00"
	location = "渋谷"
	title = "TAROUふぇすてぃばる"
	act = "山田太郎　田中たろう　TaroSakamoto"
	info = "20:00 START"
	return
}

func (bot *LineConf) CathID(req *http.Request) (userId string) {
	events, err := bot.Bot.ParseRequest(req)
	if err != nil {
		fmt.Println(err, "CathID")
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			userId = event.Source.UserID
		}
	}
	return
}

func (bot *LineConf) TestFunc(ctx context.Context, req *http.Request) (a string) {

	events, err := bot.Bot.ParseRequest(req)
	if err != nil {
		fmt.Println("ParseReq", err)
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				a = message.Text
				fmt.Println(a)

			case *linebot.StickerMessage:
				// bot.MsgReply(stickerReply)

			case *linebot.ImageMessage:
				// bot.MsgReply(imageReply)
			}
		} else {
			fmt.Println("EventTypeが違う")
		}
	}
	return

}
