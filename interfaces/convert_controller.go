package interfaces

import (
	"github.com/tmkshy1908/Portfolio/domain"
)

type ConvertController struct {
	Converter CommonConverter
}

type CommonConverter interface {
	ToSampleResponseData([]*domain.Schedule) []*domain.Schedule
}

func (cc *ConvertController) ToSampleResponseData(schedule []*domain.Schedule) (items []*domain.Schedule) {
	items = make([]*domain.Schedule, len(schedule))
	for i, s := range schedule {
		items[i] = &domain.Schedule{
			ID:       s.ID,
			Day:      s.Day,
			Contents: s.Contents,
		}
	}
	return
}
