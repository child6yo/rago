package kafka

import "github.com/IBM/sarama"

func configSarama() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V4_0_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = false

	return config
}
