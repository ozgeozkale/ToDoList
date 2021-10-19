package Models

import "time"

type TaskModel struct {
	Id          int
	Title       string
	Description string
	Category    string
	Progress    string
	Deadline    time.Time
	Priority    string
	CreatedTime time.Time
	UpdatedTime time.Time
}

type InputModel struct {
	Title       string
	Description string
	Category    string
	Progress    string
	Deadline    time.Time
	Priority    string
}

type IdModel struct {
	Id int `param:"Id" query:"Id" header:"Id" form:"Id" json:"Id" xml:"Id"`
}
