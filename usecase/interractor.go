package usecase

import (
	"context"

	"github.com/tmkshy1908/Portfolio/domain"
)

type CommonRepository interface {
	Find(ctx context.Context) (resp []*domain.Schedule, err error)
	Add(ctx context.Context, day string, contents string)
	Update(ctx context.Context, day string, contents string)
	Delete(ctx context.Context, day string)
	DivideEvent(ctx context.Context) (msg string)
	CallReply(msg string)
}
