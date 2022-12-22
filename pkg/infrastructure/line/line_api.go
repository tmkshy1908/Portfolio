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
}

func NewClient() (lh Client, err error) {
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
