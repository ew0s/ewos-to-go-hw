package generator

import (
	"context"
	"math/rand"
	"messageBrokers4/internal/domain/entities"
	"time"
)

type Prices struct {
	factor  float64
	delay   time.Duration
	tickers []string
}

func NewPrices(cfg Config) *Prices {
	return &Prices{
		factor:  cfg.Factor,
		delay:   cfg.Delay,
		tickers: cfg.Tickers,
	}
}

func (g *Prices) GeneratePrices(ctx context.Context) <-chan entities.Price {
	prices := make(chan entities.Price)

	startTime := time.Now()
	go func() {
		defer close(prices)

		ticker := time.NewTicker(g.delay)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				currentTime := time.Now()
				delta := float64(currentTime.Unix() - startTime.Unix())
				ts := time.Unix(int64(float64(currentTime.Unix())+delta*g.factor), 0)

				for idx, ticker := range g.tickers {
					minVal := float64((idx + 1) * 100)
					maxVal := minVal + 20
					prices <- entities.Price{
						Ticker: ticker,
						Value:  minVal + rand.Float64()*(maxVal-minVal),
						TS:     ts,
					}
				}
			}
		}
	}()

	return prices
}
