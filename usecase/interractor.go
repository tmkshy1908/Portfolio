package usecase

import (
	"context"

	"github.com/tmkshy1908/Portfolio/domain"
)

type CommonRepository interface {
	Find(ctx context.Context) (resp []*domain.Schedule, err error)
}
