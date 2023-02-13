package domain

import "time"

type Schedule struct {
	ID  int
	Day string
}

type Contents struct {
	ID           int
	Contents_Day time.Time
	// Contents_Day string
	Location   string
	EventTitle string
	Act        string
	OtherInfo  string
}

type Users struct {
	User_Id   string
	Condition int
}
