package interfaces

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

type Contents struct {
	// ID int
	Contents_Day time.Time
	// Contents_Day string
	Location   string
	EventTitle string
	Act        string
	OtherInfo  string
}

type dbSettings struct {
	User     string
	Password string
	Database string
}

var DB *sql.DB

func testHandler() {
	conf := dbSettings{
		User:     "yamadatarou",
		Database: "test_bot",
		Password: "1234",
	}

	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable sslmode=disable", conf.User, conf.Database, conf.Password)
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("DB Connected.")
	}
}

func Test_Xxx(t *testing.T) {
	testHandler()
	str := "2020/03/19"
	layout := "2006/01/02"
	tt, _ := time.Parse(layout, str)
	fmt.Printf("%T\n", tt)
	contents := &Contents{Contents_Day: tt, Location: "池袋", EventTitle: "イベントイベント", Act: "吉田ヨシダ　田中たなか　岡田okada", OtherInfo: "19:00 end"}
	fmt.Println(contents)
	Add(contents)
	Find()
	Update(contents)
	Delete(contents)
}

const (
	SELECT_TEST string = "select * from test_contents;"
	INSERT_TEST string = "insert into test_contents (contents_day,location,title,act,info) values ($1,$2,$3,$4,$5)"
	UPDATE_TEST string = "update test_contents set (contents_day, location, title, act,info) values($1,$2,$3,$4,$5) where contents_day = $1"
	DELETE_TEST string = "delete from test_contents where contents_day = $1)"
	// DAY_CHECK       string = "select * from test where day = TO_DATE($1, 'YY-MM-DD HH24:MI:SS')"
	// USER_CHECK      string = "select count(user_id) from users where user_id = '%s'"
	// INSERT_USERS    string = "insert into users (user_id, condition) values('%s',%b)"
	// TEST_CHECK      string = "select * from %s where %s = '%s'"
)

func Add(contents *Contents) {
	_, err := DB.Exec(INSERT_TEST, contents.Contents_Day, contents.Location, contents.EventTitle, contents.Act, contents.OtherInfo)
	if err != nil {
		fmt.Println(err, "Execエラー")
	}
}

func Find() {
	rows, err := DB.Query(SELECT_TEST)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	contents := make([]*Contents, 0)

	for rows.Next() {
		contentsTable := Contents{}
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
	fmt.Println(contents)
}

func Update(contents *Contents) {
	_, err := DB.Exec(UPDATE_TEST, &contents.Contents_Day, contents.EventTitle, contents.Location, contents.Act, contents.OtherInfo)
	if err != nil {
		fmt.Println(err, "Updateエラー")
	}
}
func Delete(contents *Contents) {
	_, err := DB.Exec(DELETE_TEST, contents.Contents_Day)
	if err != nil {
		fmt.Println(err, "Deleteエラー")
	}
}
