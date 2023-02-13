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
	DivideEvent(context.Context, *http.Request) (string, string)
	CallReply(string, string)
	WaitMsg(context.Context) (*domain.Contents, error)
	UserCheck(context.Context, string) bool
	StartUser(context.Context, string)
	EndUser(context.Context, string)
	TestTest(context.Context, *http.Request)
}
