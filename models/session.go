package models

type SessionSlice struct {
	Start time.Time `db:"start" json:"start,omitempty"`
	End   time.Time `db:"end" json:"end,omitempty"`
	Habit Habit     `db:"habit" json:"habit,omitempty"`
}

type Session struct {
	Start  time.Time       `db:"start" json:"start,omitempty"`
	End    time.Time       `db:"end" json:"end,omitempty"`
	Slices []*SessionSlice `db:"-" json:"slices"`
}
