package usecase

import (
	"context"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tmkshy1908/Portfolio/domain"
)

type CommonRepository interface {
	Find(ctx context.Context) (resp []*domain.Schedule, err error)
	DivideEvent(ctx context.Context) (e *linebot.Event)
}
