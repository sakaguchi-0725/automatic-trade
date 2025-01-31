package testutil

import "time"

var TZ = time.FixedZone("Asia/Tokyo", 9*60*60)

type FixedDateTimeFactory struct {
	year  int
	month time.Month
	day   int
}

func NewFixedDateTimeFactory(year int, month time.Month, day int) FixedDateTimeFactory {
	return FixedDateTimeFactory{year, month, day}
}

func (factory FixedDateTimeFactory) At(hour, minute int) time.Time {
	return time.Date(factory.year, factory.month, factory.day, hour, minute, 0, 0, TZ)
}
