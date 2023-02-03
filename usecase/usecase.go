package usecase

import (
	"context"
	"encoding/json"
	"fmt"
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
	// if strings.Contains(msg, "編集") {
	// msg = "編集モードです"
	// i.CommonRepository.CallReply(msg)
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
		i.CommonRepository.CallReply(msg, userId)
		resp, err := i.CommonRepository.WaitMsg(ctx)
		if err != nil {
			fmt.Println(err, "WaitMsgエラー")
		}
		err = i.CommonRepository.Add(ctx, resp)
		if err != nil {
			fmt.Println(err)
		}
		// i.CommonRepository.CallReply(dayRequestMsg)
		// day := i.CommonRepository.DivideEvent(ctx)
		// i.CommonRepository.CallReply(contentsRequestMsg)
		// contents := i.CommonRepository.DivideEvent(ctx)
		// i.CommonRepository.Add(ctx, day, contents)

	} else if strings.Contains(msg, "更新") {
		resp, err := i.CommonRepository.WaitMsg(ctx)
		if err != nil {
			fmt.Println(err)
		}
		err = i.CommonRepository.Update(ctx, resp)
		if err != nil {
			fmt.Println(err)
		}
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
		i.CommonRepository.CallReply(msg, userId)

	} else if strings.Contains(msg, "test") {
		i.CommonRepository.TestTest(ctx, req)
		i.CommonRepository.UserCheck(ctx, userId)
	}
	// }

}
