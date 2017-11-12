package models

type Habit struct {
	Id          string `db:"id" json:"id,omitempty"`
	Title       string `db:"title" json:"title,omitempty"`
	AvgDuration uint32 `db:"avgDuration" json:"avgDuration,omitempty"`
	Deadline    string `db:"Deadline" json:"deadline,omitempty"`
	During      uint32 `db:"during" json:"during,omitempty"`
	Every       uint32 `db:"every" json:"every,omitempty"`
	EveryUnit   string `db:"everyUnit" json:"everyUnit,omitempty"`
}
