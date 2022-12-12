package interfaces

import (
	"context"
	"fmt"

	"github.com/tmkshy1908/Portfolio/domain"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
)

type CommonRepository struct {
	DB db.SqlHandler
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

// func (r *CommonRepository) Create(ctx context.Context) (schedule []*domain.Schedule, err error){
// 	_, err = r.DB.Exec(ctx, INSERT_SCHEDULE)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// defer Close()

// }
