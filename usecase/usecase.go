package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tmkshy1908/Portfolio/domain"
)

type CommonInteractor struct {
	CommonRepository CommonRepository
}

func (i *CommonInteractor) UseCaseSampleRepository(ctx context.Context) (resp []*domain.Schedule, err error) {
	msg := i.CommonRepository.DivideEvent(ctx)
	fmt.Println(msg, "Usecase")
	if strings.Contains(msg, "取得") {
		i.CommonRepository.CallReply(msg)
		resp, err = i.CommonRepository.Find(ctx)
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
	}

	return
}

// func (i *CommonInteractor) UseCaseLineRepository(ctx context.Context) {
// 	fmt.Println("LineRep")
// i.CommonRepository.DivideEvent(ctx)
// event := i.CommonRepository.DivideEvent(ctx)
// msg := event.Message
// fmt.Println(msg, "msgだよ")
// }
