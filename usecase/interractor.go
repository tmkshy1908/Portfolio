package usecase

import (
	"context"

	"github.com/tmkshy1908/Portfolio/domain"
)

type CommonRepository interface {
	Find(context.Context) ([]*domain.Contents, error)
	Add(context.Context, *domain.Contents) error
	Update(context.Context, string, string) error
	Delete(context.Context, string) error
	DivideEvent(context.Context) string
	CallReply(string)
	WaitMsg(context.Context) (*domain.Contents, error)
}
