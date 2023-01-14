package domain

import (
	"time"
)

type Schedule struct {
	ID  int
	Day string
}

type Contents struct {
	ID           int
	Contents_Day time.Time
	Location     string
	EventTitle   string
	Act          string
	OtherInfo    string
}
