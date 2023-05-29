package interfaces

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/tmkshy1908/Portfolio/domain"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/line"
)

type CommonRepository struct {
	DB  db.SqlHandler
	Bot line.Client
}

const (
	SELECT_CONTENTS string = "select contents_day, location, event_title, act, other_info from contents;"
	INSERT_CONTENTS string = "insert into  contents (contents_day, location, event_title, act, other_info) values($1,$2,$3,$4,$5)"
	INSERT_SCHEDULE string = "insert into schedule (day) values ($1)"
	UPDATE_CONTENTS string = "update contents set (contents_day, location, event_title, act, other_info) = ($1,$2,$3,$4,$5) where contents_day = $1"
	DELETE_SCHEDULE string = "delete from schedule where day = $1"
	DELETE_CONTENTS string = "delete from contents where contents_day = $1"
	USER_CHECK      string = "select count(user_id) from users where user_id = $1"
	CONDITION_CHECK string = "select condition from users where user_id = $1"
	INSERT_USERS    string = "insert into users (user_id, condition) values($1,$2)"
	DELETE_USERS    string = "delete from users where user_id = $1"
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
	err = r.DB.ExecWithTx(func(*sql.Tx) error {
		_, err = r.DB.Exec(ctx, INSERT_SCHEDULE, contents.Contents_Day)
		if err != nil {
			fmt.Println(err, "Execエラー")
			return err
		}
		_, err = r.DB.Exec(ctx, INSERT_CONTENTS, contents.Contents_Day, contents.Location, contents.EventTitle, contents.Act, contents.OtherInfo)
		if err != nil {
			fmt.Println(err, "Execエラー")
			return err
		}
		return err
	})
	return
}

func (r *CommonRepository) Update(ctx context.Context, contents *domain.Contents) (err error) {
	_, err = r.DB.Exec(ctx, UPDATE_CONTENTS, contents.Contents_Day, contents.EventTitle, contents.Location, contents.Act, contents.OtherInfo)
	if err != nil {
		fmt.Println(err, "Updateエラー")
		return err
	}
	return
}

func (r *CommonRepository) Delete(ctx context.Context, contents *domain.Contents) (err error) {
	err = r.DB.ExecWithTx(func(*sql.Tx) error {
		_, err = r.DB.Exec(ctx, DELETE_CONTENTS, contents.Contents_Day)
		if err != nil {
			fmt.Println(err, "Deleteエラー")
			return err
		}
		_, err = r.DB.Exec(ctx, DELETE_SCHEDULE, contents.Contents_Day)
		if err != nil {
			fmt.Println(err, "Deleteエラー")
			return err
		}
		return err
	})
	return
}

func (r *CommonRepository) DivideEvent(ctx context.Context, req *http.Request) (msg string, userId string) {
	msg, userId = r.Bot.CathEvents(ctx, req)
	return
}

func (r *CommonRepository) CallReply(msg string, userId string) {
	r.Bot.MsgReply(msg, userId)
}

func (r *CommonRepository) WaitMsg(ctx context.Context) (contents *domain.Contents, err error) {
	day, location, title, act, info := r.Bot.WaitEvents(ctx)
	contents = &domain.Contents{Contents_Day: day, Location: location, EventTitle: title, Act: act, OtherInfo: info}
	return
}

func (r *CommonRepository) UserCheck(ctx context.Context, userId string) bool {
	var i int
	// UserIdが登録されているかのチェック i=0 は登録なし i=1 は登録ずみ
	err := r.DB.QueryRow(ctx, USER_CHECK, userId).Scan(&i)
	if err != nil {
		fmt.Println(err, "idチェック")
	}
	if i == 0 {
		return false
	} else {
		return true
	}
	// return true
}

func (r *CommonRepository) StartUser(ctx context.Context, userId string) {
	var i int
	// UserIdが登録されているかのチェック i=0 は登録なし i=1 は登録ずみ
	err := r.DB.ExecWithTx(func(*sql.Tx) error {
		err := r.DB.QueryRow(ctx, USER_CHECK, userId).Scan(&i)
		if err != nil {
			fmt.Println(err, "StartUser:QueryRow")
		}
		if i == 0 {
			_, err := r.DB.Exec(ctx, INSERT_USERS, userId, 0)
			if err != nil {
				fmt.Println(err, "StartUser:Exec")
			}
		}
		return err
	})
	fmt.Println(err)
}

func (r *CommonRepository) EndUser(ctx context.Context, userId string) {
	_, err := r.DB.Exec(ctx, DELETE_USERS, userId)
	if err != nil {
		fmt.Println(err, "EndUser")
	}
}

func (r *CommonRepository) ConditionCheck(ctx context.Context, userId string) (cNumber int) {
	err := r.DB.QueryRow(ctx, USER_CHECK, userId).Scan(&cNumber)
	if err != nil {
		fmt.Println(err, "クエリ")
	}
	fmt.Println(cNumber)
	return cNumber
}
