package interfaces

import (
	"context"
	"fmt"

	"github.com/tmkshy1908/Portfolio/domain"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/line"
)

type CommonRepository struct {
	DB  db.SqlHandler
	Bot line.Client
}

const (
	SELECT_SCHEDULE string = "select id, day, contents from schedule;"
	INSERT_SCHEDULE string = "insert into schedule (day, contents) values($1,$2)"
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

func (r *CommonRepository) Add(ctx context.Context, day string, contents string) {
	// schedule := make([]*domain.Schedule, 0)
	values := fmt.Sprintf("insert into schedule (day, contents) values(%s,'%s')", day, contents)
	_, err := r.DB.Exec(ctx, values)
	if err != nil {
		fmt.Println(err, "Execエラー")
		return
	}
}

func (r *CommonRepository) DivideEvent(ctx context.Context) (msg string) {
	msg = r.Bot.CathEvents(ctx)
	fmt.Println("DivideEvent", msg)

	return
}

func (r *CommonRepository) CallReply(msg string) {
	r.Bot.MsgReply(msg)
}
