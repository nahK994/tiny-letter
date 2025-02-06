package mq_producer

import (
	"encoding/json"
	"fmt"
	"tiny-letter/subscription/pkg/app"

	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer(config *app.MQ_config) (*Producer, error) {
	producer, err := connectProducer(config)
	if err != nil {
		return nil, err
	}
	defer producer.Close()
	return &Producer{producer: producer}, nil
}

func connectProducer(config *app.MQ_config) (sarama.SyncProducer, error) {
	broker := fmt.Sprintf("%s:%d", config.Domain, config.Port)

	mqConfig := sarama.NewConfig()
	mqConfig.Producer.Return.Successes = config.IsProducerReturnSuccess
	mqConfig.Producer.RequiredAcks = sarama.WaitForAll
	mqConfig.Producer.Retry.Max = config.NumberOfRetry

	return sarama.NewSyncProducer([]string{broker}, mqConfig)
}

func (p *Producer) Push(data interface{}) error {
	message, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: app.GetConfig().MQ.Topic,
		Value: sarama.StringEncoder(message),
	}
	_, _, err = p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
