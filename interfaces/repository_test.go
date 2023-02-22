package interfaces_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/tmkshy1908/Portfolio/domain"
	"github.com/tmkshy1908/Portfolio/usecase"
)

type testInteractor struct {
	usecase.CommonRepository
}

func (i testInteractor) Test_Add(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()
	str := "2020/03/19"
	layout := "2006/01/02"
	tt, _ := time.Parse(layout, str)
	contents := &domain.Contents{Contents_Day: tt, Location: "池袋", EventTitle: "イベントイベント", Act: "吉田ヨシダ　田中たなか　岡田okada", OtherInfo: "19:00 end"}
	err := i.Add(ctx, contents)
	if err != nil {
		fmt.Println(err)
	}
}

// func (i CommonRepository)TestFind(t *testing.T) {
// 	i.Find()
// }

// func (i testInteractor) TestXxx(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
// 	defer cancel()
// 	str := "2020/03/19"
// 	layout := "2006/01/02"
// 	tt, _ := time.Parse(layout, str)
// 	contents := &domain.Contents{Contents_Day: tt, Location: "池袋", EventTitle: "イベントイベント", Act: "吉田ヨシダ　田中たなか　岡田okada", OtherInfo: "19:00 end"}
// 	i.CommonRepository.Add(ctx, contents)
// 	i.CommonRepository.Find(ctx)
// 	contents = &domain.Contents{Contents_Day: tt, Location: "池袋", EventTitle: "イベントイベント", Act: "アップデートしました", OtherInfo: "19:00 end"}
// 	i.CommonRepository.Update(ctx, contents)
// 	i.CommonRepository.Find(ctx)
// 	i.CommonRepository.Delete(ctx, contents)
// }
