package generator

import "time"

type Config struct {
	Factor  float64
	Delay   time.Duration
	Tickers []string
}
