package timeservice

import "fmt"

type TimeService struct {
	TotalTicks int
}

func NewTimeSerice() *TimeService {
	return &TimeService{
		TotalTicks: 0,
	}
}

func (t *TimeService) GetPrettySeconds() string {
	fmt.Println(t.TotalTicks / 60)
	return ""
}
