package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type CommonInteractor struct {
	CommonRepository CommonRepository
}

func (i *CommonInteractor) DivideMessage(ctx context.Context) {
	msg := i.CommonRepository.DivideEvent(ctx)
	if strings.Contains(msg, "取得") {
		i.CommonRepository.CallReply(msg)
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
		i.CommonRepository.CallReply(string(out))
	} else if strings.Contains(msg, "作成") {
		i.CommonRepository.CallReply(msg)
		day := "20221212"
		contents := "こんにちは"
		i.CommonRepository.Add(ctx, day, contents)
	}

}
