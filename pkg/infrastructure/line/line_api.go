package line

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

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
	WaitEvents(context.Context) (string, string)
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

func (bot *LineConf) WaitEvents(ctx context.Context) (day string, contents string) {
	a := "2022年04月01日" + "T09:00:00+09:00"
	parsedTime, _ := time.Parse("2006年01月02日T15:04:05Z07:00", a)
	fmt.Println(parsedTime)
	contents = "こんにちは"
	return
}

// func (bot *LineConf)
