package entities

import (
	"time"
)

type Price struct {
	Ticker string
	Value  float64
	TS     time.Time
}
