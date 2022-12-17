package api

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

type LineHandller interface {
	CathEvents(ctx context.Context) (msg string)
	MsgReply(msg string)
	// Hoge(event *linebot.Event) (e *linebot.Event)
}

func NewLineHandller() (lh LineHandller, err error) {
	bot, err := linebot.New(
		"4a7eaa800c243575a028db8438842246",
		"P5L9UuMlMuG6sRbGgC0N/rGfICCAZ4P0ixLf7hgomVVyqxHvD5G4ZHNqu7IxpkpYut2LJ5NJ1qgKtCBveIIx4MZGOzuR6ldFGC33TBOXktYbHGhHY7bwQuolurMpN5YW/enP8ZNWUdBjE7PeqGEOswdB04t89/1O/w1cDnyilFU=",
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
	} else {
		fmt.Println("CathEvents")
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

func (bot *LineConf) MsgReply(msg string) {
	replyMessage := linebot.NewTextMessage(msg)
	bot.Bot.BroadcastMessage(replyMessage).Do()
}

const (
	stickerReply string = "いいスタンプだね！"
	imageReply   string = "素敵な写真だね！"
)

// func (bot *LineConf) Hoge(event *linebot.Event) (e *linebot.Event) {
// 	switch event.Message.(type) {
// 	case *linebot.TextMessage:
// 		e = event
// 	case *linebot.StickerMessage:
// 		bot.MsgReply(stickerReply)

// 	case *linebot.ImageMessage:
// 		bot.MsgReply(imageReply)
// 	}
// 	return
// }

// func (l *LineConf) CathEvents(c echo.Context) (e *linebot.Event) {
// 	events, err := linebot.ParseRequest(c.Request())
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for _, event := range events {
// 		if event.Type == linebot.EventTypeMessage {
// 			switch message := event.Message.(type) {

// 			case *linebot.TextMessage:
// 				replyMessage := message.Text

// 			}
// 		}
// 	}
// }
