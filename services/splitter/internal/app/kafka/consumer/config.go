package consumer

import (
	"time"

	"github.com/IBM/sarama"
)

func configSarama() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Return.Errors = true

	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	config.Consumer.Group.Session.Timeout = 6 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 2 * time.Second

	return config
}
