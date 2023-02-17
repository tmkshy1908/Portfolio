package interfaces

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/tmkshy1908/Portfolio/domain"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/line"
	"github.com/tmkshy1908/Portfolio/usecase"
)

type dbSettings struct {
	User     string
	Password string
	Database string
}

// var DB *sql.DB
type testConf struct {
	Conn *db.SqlConf
}

func TestHandler() (h db.SqlHandler, err error) {
	conf := dbSettings{
		User:     "yamadatarou",
		Database: "test_bot",
		Password: "1234",
	}

	// var err error
	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable sslmode=disable", conf.User, conf.Database, conf.Password)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("DB Connected.")
	}
	h = &testConf{Conn.Conn: db}
	return
}

//

func TestXxx(t *testing.T) {
	bot, err := line.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	db, err := TestHandler()
	if err != nil {
		fmt.Println(err)
	}

	r := NewController(db, bot)
}

type testInteractor struct {
	usecase.CommonRepository
}

func (i testInteractor) TestXxx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	str := "2020/03/19"
	layout := "2006/01/02"
	tt, _ := time.Parse(layout, str)
	contents := &domain.Contents{Contents_Day: tt, Location: "池袋", EventTitle: "イベントイベント", Act: "吉田ヨシダ　田中たなか　岡田okada", OtherInfo: "19:00 end"}
	i.CommonRepository.Add(ctx, contents)
	i.CommonRepository.Find(ctx)
	contents = &domain.Contents{Contents_Day: tt, Location: "池袋", EventTitle: "イベントイベント", Act: "アップデートしました", OtherInfo: "19:00 end"}
	i.CommonRepository.Update(ctx, contents)
	i.CommonRepository.Find(ctx)
	i.CommonRepository.Delete(ctx, contents)
}
