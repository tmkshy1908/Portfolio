package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type CommonInteractor struct {
	CommonRepository CommonRepository
}

// const (
// 	dayRequestMsg      string = "日付を入力してください"
// 	contentsRequestMsg string = "内容を入力してください"
// )

func (i *CommonInteractor) DivideMessage(ctx context.Context, req *http.Request) {
	msg, userId := i.CommonRepository.DivideEvent(ctx, req)

	if strings.Contains(msg, "編集") {
		msg = "編集モード"
		i.CommonRepository.StartUser(ctx, userId)
		i.CommonRepository.CallReply(msg, userId)
	} else if strings.Contains(msg, "終了") {
		msg = "終了しました"
		i.CommonRepository.EndUser(ctx, userId)
		i.CommonRepository.CallReply(msg, userId)
	}

	if i.CommonRepository.UserCheck(ctx, userId) {
		if strings.Contains(msg, "取得") {
			i.CommonRepository.CallReply(msg, userId)
			resp, err := i.CommonRepository.Find(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}
			a := &resp
			out, err := json.Marshal(a)
			if err != nil {
				fmt.Println("marshal err")
			}

			i.CommonRepository.CallReply(string(out), userId)
			fmt.Printf("%T\n", out)

		} else if strings.Contains(msg, "作成") {
			resp, err := i.CommonRepository.WaitMsg(ctx)
			if err != nil {
				fmt.Println(err, "WaitMsgエラー")
			}
			err = i.CommonRepository.Add(ctx, resp)
			if err != nil {
				fmt.Println(err)
			}
			msg = "作成しました"
			i.CommonRepository.CallReply(msg, userId)

		} else if strings.Contains(msg, "更新") {
			resp, err := i.CommonRepository.WaitMsg(ctx)
			if err != nil {
				fmt.Println(err)
			}
			err = i.CommonRepository.Update(ctx, resp)
			if err != nil {
				fmt.Println(err)
			}
			msg = "更新しました"
			i.CommonRepository.CallReply(msg, userId)

		} else if strings.Contains(msg, "削除") {
			resp, err := i.CommonRepository.WaitMsg(ctx)
			if err != nil {
				fmt.Println(err)
			}
			err = i.CommonRepository.Delete(ctx, resp)
			if err != nil {
				fmt.Println(err)
			}
			msg = "削除しました"
			i.CommonRepository.CallReply(msg, userId)

		} else if strings.Contains(msg, "test") {
			resp, err := http.Get("https://3517-240d-0-4b77-6a00-78b6-f3a5-c59c-d96f.jp.ngrok.io/line")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(resp.StatusCode)
			fmt.Println(resp.Header["Content-Type"])
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err, "ioutil")
			}
			fmt.Print(string(body))

		}
	}
	// i.CommonRepository.CallReply(msg, userId)
}
