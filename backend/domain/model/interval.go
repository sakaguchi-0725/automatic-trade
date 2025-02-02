package model

type Interval string

const (
	Min5  Interval = "5"
	Min15 Interval = "15"
	Hour1 Interval = "60"
	Hour4 Interval = "240"
	Day   Interval = "D"
	Month Interval = "M"
	Week  Interval = "W"
)
