package timeservice

import (
	"strconv"
)

type TimeService struct {
	TotalTicks int
}

func NewTimeSerice() *TimeService {
	return &TimeService{
		TotalTicks: 0,
	}
}

func (t *TimeService) GetPrettySeconds() string {
	var time = t.TotalTicks / 60
	return strconv.Itoa(time)
}
