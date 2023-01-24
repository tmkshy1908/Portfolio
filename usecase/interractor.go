package usecase

import (
	"context"
	"net/http"

	"github.com/tmkshy1908/Portfolio/domain"
)

type CommonRepository interface {
	Find(context.Context) ([]*domain.Contents, error)
	Add(context.Context, *domain.Contents) error
	Update(context.Context, *domain.Contents) error
	Delete(context.Context, *domain.Contents) error
	DivideEvent(context.Context, *http.Request) string
	CallReply(string)
	WaitMsg(context.Context) (*domain.Contents, error)
	TestTest(context.Context, *http.Request)
}
