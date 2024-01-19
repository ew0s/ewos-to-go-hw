package producer

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/IBM/sarama"

	"messageBrokers4/internal/domain/generator"
)

var tickers = []string{"AAPL", "SBER", "NVDA", "TSLA"}

type Prices struct {
	log      *log.Logger
	producer sarama.SyncProducer
}

func NewPrices() *Prices {
	// TODO: напиши логику конструктора
	panic("implement me")
}

func (p *Prices) ProducePrices(ctx context.Context) {
	pg := generator.NewPrices(generator.Config{
		Factor:  10,
		Delay:   time.Millisecond * 500,
		Tickers: tickers,
	})

	p.log.Info("start prices generator...")
	_ = pg.GeneratePrices(ctx)

	// TODO: напиши логику продюсера
	panic("implement me")
}
