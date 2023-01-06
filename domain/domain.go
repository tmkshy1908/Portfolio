package domain

// type Schedule struct {
// 	ID       int
// 	Day      string
// 	Contents string
// }

type Schedule struct {
	ID           int
	Day          string
	ContentsList []*Contents
}

type Contents struct {
	ID           int
	Contents_Day int
	Location     string
	EventTile    string
	Act          string
	OtherInfo    string
}
