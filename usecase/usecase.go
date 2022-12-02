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
	resp, err = i.CommonRepository.Find(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
