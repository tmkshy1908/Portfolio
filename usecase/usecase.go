package usecase

import (
	"context"
	"fmt"

	"github.com/tmkshy1908/Portfolio/domain"
)

type CommonInteractor struct {
	CommonRepository CommonRepository
}

func (i *CommonInteractor) UseCaseSampleRepository(ctx context.Context) (resp []*domain.Schedule, err error) {
	// event := i.CommonRepository.DivideEvent(ctx)
	// msg := event.Message
	// fmt.Println(msg, "msgだよ")
	msg := i.CommonRepository.DivideEvent(ctx)
	fmt.Println(msg, "Usecase")
	resp, err = i.CommonRepository.Find(ctx)
	if err != nil {
		fmt.Println(err)
		return
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
