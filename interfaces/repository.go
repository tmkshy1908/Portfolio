package interfaces

import (
	"context"
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tmkshy1908/Portfolio/domain"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/api"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
)

type CommonRepository struct {
	DB  db.SqlHandler
	Bot api.LineHandller
}

const (
	SELECT_SCHEDULE string = "select id, day, contents from schedule;"
	INSERT_SCHEDULE string = "insert into schedule (id, day, contents) values($1,$2,$3)"
)

func (r *CommonRepository) Find(ctx context.Context) (schedule []*domain.Schedule, err error) {
	rows, err := r.DB.Query(ctx, SELECT_SCHEDULE)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	schedule = make([]*domain.Schedule, 0)

	for rows.Next() {
		scheduleTable := domain.Schedule{}
		if err = rows.Scan(
			&scheduleTable.ID,
			&scheduleTable.Day,
			&scheduleTable.Contents,
		); err != nil {
			fmt.Println(err)
			return
		}
		schedule = append(schedule, &scheduleTable)
	}
	return
}

func (r *CommonRepository) DivideEvent(ctx context.Context) (e *linebot.Event) {
	event := r.Bot.CathEvents(ctx)
	fmt.Println("DivideEvent", event)

	return
}

// func (r *CommonRepository) ConversionReply(resp []*domain.Schedule) {
// 	// 文字列に直す処理
// 	// r.Bot.MsgReply(resp)
// }
