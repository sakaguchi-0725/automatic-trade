package model

import (
	"errors"
	"time"
)

type Rate struct {
	DateTime time.Time
	Price    float64
}

type Rates []Rate

func (r Rates) Latest() (Rate, error) {
	if len(r) == 0 {
		return Rate{}, errors.New("no rates available")
	}

	latest := r[0]
	for _, rate := range r {
		if rate.DateTime.After(latest.DateTime) {
			latest = rate
		}
	}

	return latest, nil
}
