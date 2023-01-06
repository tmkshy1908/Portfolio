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
	// SELECT_SCHEDULE string = "select * from schedule;"
	SELECT_CONTENTS string = "select * from contents;"
	INSERT_SCHEDULE string = "insert into schedule (day, contents) values(%s,'%s')"
	UPDATE_SCHEDULE string = "update schedule set day = '%s', contents = '%s' where day = '%s'"
	DELETE_SCHEDULE string = "delete from schedule where day = '%s'"
)

func (r *CommonRepository) Find(ctx context.Context) (contents []*domain.Contents, err error) {
	rows, err := r.DB.Query(ctx, SELECT_CONTENTS)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	contents = make([]*domain.Contents, 0)

	for rows.Next() {
		contentsTable := domain.Contents{}
		if err = rows.Scan(
			&contentsTable.ID,
			&contentsTable.Contents_Day,
			&contentsTable.Location,
			&contentsTable.EventTile,
			&contentsTable.Act,
			&contentsTable.OtherInfo,
		); err != nil {
			fmt.Println(err)
			return
		}
		contents = append(contents, &contentsTable)
	}
	return
}

func (r *CommonRepository) Add(ctx context.Context, day string, contents string) (err error) {
	values := fmt.Sprintf(INSERT_SCHEDULE, day, contents)
	_, err = r.DB.Exec(ctx, values)
	if err != nil {
		// fmt.Println(err, "Execエラー")
		return err
	}
	return
}

func (r *CommonRepository) Update(ctx context.Context, day string, contents string) (err error) {
	values := fmt.Sprintf(UPDATE_SCHEDULE, day, contents, day)
	_, err = r.DB.Exec(ctx, values)
	if err != nil {
		// fmt.Println(err, "Updateエラー")
		return err
	}
	return
}

func (r *CommonRepository) Delete(ctx context.Context, day string) (err error) {
	values := fmt.Sprintf(DELETE_SCHEDULE, day)
	_, err = r.DB.Exec(ctx, values)
	if err != nil {
		// fmt.Println(err, "Deleteエラー")
		return err
	}
	return
}

func (r *CommonRepository) DivideEvent(ctx context.Context) (msg string) {
	msg = r.Bot.CathEvents(ctx)
	return
}

func (r *CommonRepository) CallReply(msg string) {
	r.Bot.MsgReply(msg)
	fmt.Printf("%T\n", msg)
}

func (r *CommonRepository) WaitMsg(ctx context.Context) (day string, contents string) {
	day, contents = r.Bot.WaitEvents(ctx)
	return
}
