package consumer

import (
	"github.com/IBM/sarama"
)

type Prices struct {
}

func (c *Prices) Setup(session sarama.ConsumerGroupSession) error {
	//TODO implement me
	panic("implement me")
}

func (c *Prices) Cleanup(session sarama.ConsumerGroupSession) error {
	//TODO implement me
	panic("implement me")
}

func (c *Prices) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	//TODO implement me
	panic("implement me")
}
