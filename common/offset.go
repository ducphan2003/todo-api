package common

import "time"

type Offset struct {
	Offset int `json:"offset" form:"offset"`
	Limit  int `json:"limit" form:"limit"`
}

type LastUpdate struct {
	FromTime *time.Time `json:"from_time" form:"from_time"`
	Limit    *int       `json:"limit" form:"limit"`
}

func (l *LastUpdate) Fill() {
	if l.FromTime == nil {
		now := time.Now()
		l.FromTime = &now
	}

	if l.Limit == nil {
		var defaultLimit = 50
		l.Limit = &defaultLimit
	}
}
