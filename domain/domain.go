package domain

import (
	"time"
)

type Schedule struct {
	ID           int
	Day          string
	ContentsList []*Contents
}

type Contents struct {
	ID           int
	Contents_Day time.Time
	Location     string
	EventTile    string
	Act          string
	OtherInfo    string
}
