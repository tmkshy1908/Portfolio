package line

import (
	"context"
	"fmt"
	"net/http"

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
	CathEvents(ctx context.Context) (msg string)
	MsgReply(msg string)
	WaitEvents(ctx context.Context) (day string, contents string)
}

func NewClient() (lh Client, err error) {
	bot, err := linebot.New(
		"qaZpE0HLQyRNQO3L1L7QBQ+hzgpwUJflVjHILt5TGgYo/ib4gdDHkjHsZNnRTiCUut2LJ5NJ1qgKtCBveIIx4MZGOzuR6ldFGC33TBOXkta1piztcg7piep4V1rCemgIIBCpAODkqQWq2d74AhKH/AdB04t89/1O/w1cDnyilFU=",
		"b4c80739cd0c020986f94977946f696e",
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
	fmt.Println(events)
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
	day = "20221212"
	contents = "こんにちは"
	return
}
