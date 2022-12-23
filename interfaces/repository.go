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
	INSERT_SCHEDULE string = "insert into schedule (day, contents) values(%s,'%s')"
	UPDATE_SCHEDULE string = "update schedule set day = '%s', contents = '%s' where day = '%s'"
	DELETE_SCHEDULE string = "delete from schedule where day = '%s'"
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
	values := fmt.Sprintf(INSERT_SCHEDULE, day, contents)
	_, err := r.DB.Exec(ctx, values)
	if err != nil {
		fmt.Println(err, "Execエラー")
		return
	}
}

func (r *CommonRepository) Update(ctx context.Context, day string, contents string) {
	values := fmt.Sprintf(UPDATE_SCHEDULE, day, contents, day)
	_, err := r.DB.Exec(ctx, values)
	if err != nil {
		fmt.Println(err, "Updateエラー")
		return
	}
}

func (r *CommonRepository) Delete(ctx context.Context, day string) {
	values := fmt.Sprintf(DELETE_SCHEDULE, day)
	_, err := r.DB.Exec(ctx, values)
	if err != nil {
		fmt.Println(err, "Deleteエラー")
		return
	}
}

func (r *CommonRepository) DivideEvent(ctx context.Context) (msg string) {
	msg = r.Bot.CathEvents(ctx)
	return
}

func (r *CommonRepository) CallReply(msg string) {
	r.Bot.MsgReply(msg)
}
