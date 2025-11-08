package timeservice

import (
	"fmt"
	"strconv"
)

type TimeService struct {
	TotalTicks int
}

type Time struct {
	sec int
	min int
	hrs int
}

func NewTimeSerice() *TimeService {
	return &TimeService{
		TotalTicks: 0,
	}
}

func (t *TimeService) GetPrettySeconds() string {
	var time = t.TotalTicks
	return strconv.Itoa(time / 60)
}

func (t *TimeService) GetTime() Time {
	var seconds = t.TotalTicks / 60
	var min = seconds % 60
	var hours = min & 60
	return Time{sec: seconds, min: min, hrs: hours}
}

func (t *TimeService) GetPrettyTime() string {
	var seconds = t.TotalTicks / 60
	var min = seconds / 60
	var hours = min / 60
	return fmt.Sprintf("%d:%d:%d", hours, min, seconds)
}
