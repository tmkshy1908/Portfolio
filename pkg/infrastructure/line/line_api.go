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
	CathEvents(context.Context) string
	MsgReply(string)
	WaitEvents(context.Context) (string, string, string, string, string)
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
	lh = &LineConf{Bot: bot}

	return
}

func (bot *LineConf) CathEvents(ctx context.Context) (msg string) {
	events, err := bot.Bot.ParseRequest(ctx.Value("request").(*http.Request))
	if err != nil {
		fmt.Println("ParseReq", err)
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msg = message.Text

			case *linebot.StickerMessage:
				bot.MsgReply(stickerReply)

			case *linebot.ImageMessage:
				bot.MsgReply(imageReply)
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

func (bot *LineConf) MsgReply(msg string) {
	replyMessage := linebot.NewTextMessage(msg)
	bot.Bot.BroadcastMessage(replyMessage).Do()
}

func (bot *LineConf) WaitEvents(ctx context.Context) (day string, location string, title string, act string, info string) {
	day = "22-04-01"
	location = "渋谷"
	title = "TAROUふぇすてぃばる"
	act = "山田太郎　田中たろう　TaroSakamoto"
	info = "20:00 START"
	return
}

// func (bot *LineConf)
