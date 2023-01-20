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
	INSERT_CONTENTS string = "insert into contents (contents_day, location, event_title, act, other_info) values(TO_DATE('%s', 'YY-MM-DD'),'%s','%s','%s','%s')"
	UPDATE_CONTENTS string = "update contents set (contents_day, location, event_title, act, other_info) values(TO_DATE('%s', 'YY-MM-DD'),'%s','%s','%s','%s') "
	DELETE_CONTENTS string = "delete from contents where contents_day = TO_DATE('%s', 'YY-MM-DD')"
	DAY_CHECK       string = "select * from test where day = TO_DATE($1, 'YY-MM-DD HH24:MI:SS')"
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
			&contentsTable.EventTitle,
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

func (r *CommonRepository) Add(ctx context.Context, contents *domain.Contents) (err error) {
	fmt.Println(&contents)
	fmt.Println(contents.Contents_Day, contents.EventTitle, contents.Location, contents.Act, contents.OtherInfo)
	// contentsTable := make([]*domain.Contents,0)saa
	value := fmt.Sprintf("insert into schedule (day) values (TO_DATE('%s', 'YY-MM-DD HH24:MI:SS'))", contents.Contents_Day)
	_, err = r.DB.Exec(ctx, value)
	if err != nil {
		fmt.Println(err, "Execエラー")
		return err
	}
	values := fmt.Sprintf(INSERT_CONTENTS, contents.Contents_Day, contents.Location, contents.EventTitle, contents.Act, contents.OtherInfo)
	_, err = r.DB.Exec(ctx, values)
	if err != nil {
		fmt.Println(err, "Execエラー")
		return err
	}
	return
}

func (r *CommonRepository) Update(ctx context.Context, contents *domain.Contents) (err error) {
	values := fmt.Sprintf(UPDATE_CONTENTS, contents.Contents_Day, contents.EventTitle, contents.Location, contents.Act, contents.OtherInfo)
	_, err = r.DB.Exec(ctx, values)
	if err != nil {
		fmt.Println(err, "Updateエラー")
		return err
	}
	return
}

func (r *CommonRepository) Delete(ctx context.Context, contents *domain.Contents) (err error) {
	values := fmt.Sprintf(DELETE_CONTENTS, contents.Contents_Day)
	_, err = r.DB.Exec(ctx, values)
	if err != nil {
		fmt.Println(err, "Deleteエラー")
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
}

func (r *CommonRepository) WaitMsg(ctx context.Context) (contents *domain.Contents, err error) {
	day, location, title, act, info := r.Bot.WaitEvents(ctx)
	// contents_day := day + " 00:00:00"
	// contents_day, _ := time.Parse("2006年01月02日T15:04:05Z07:00", a)
	fmt.Println(day)
	contents = &domain.Contents{Contents_Day: day, Location: location, EventTitle: title, Act: act, OtherInfo: info}
	// contents = append(contents, &contentsTable)
	return
}

// func (r *CommonRepository) dayCheck(ctx context.Context, day string) {
// 	values := fmt.Sprintf(DAY_CHECK, day)
// 	_, err := r.DB.Exec(ctx,values)
// 	if err != nil {
// 		fmt.Println("Create Execエラー:", err)
// 		t = false
// 		return
// 	}
// 	t = true
// }

func (r *CommonRepository) TestTest(ctx context.Context) {
	for i := 0; i < 5; i++ {
		r.Bot.TestFunc(ctx)
	}
	fmt.Println("ループ処理終わり")
}
