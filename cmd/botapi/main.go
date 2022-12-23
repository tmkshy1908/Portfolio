package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/db"
	"github.com/tmkshy1908/Portfolio/pkg/infrastructure/line"
)

func main() {
	loadEnv()
	bot, err := line.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	db, err := db.NewHandler()
	if err != nil {
		fmt.Println(err)
	}
	infrastructure.NewServer(db, bot)

}

func loadEnv() {
	// ここで.envファイル全体を読み込みます。
	// この読み込み処理がないと、個々の環境変数が取得出来ません。
	// 読み込めなかったら err にエラーが入ります。
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// .envの SAMPLE_MESSAGEを取得して、messageに代入します。
	message := os.Getenv("SAMPLE_MESSAGE")

	fmt.Println(message)
}
