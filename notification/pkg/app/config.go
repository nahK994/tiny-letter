package app

import constant "tiny-letter/notification/pkg/constants"

const (
	localhost = "localhost"
)

type CommConfig struct {
	Port   int
	Domain string
}

type QueueConfig struct {
	Topic string
}

type Consumer struct {
	Confirmation          QueueConfig
	Publication           QueueConfig
	IsConsumerReturnError bool
}

type Producer struct {
	Confirmation            QueueConfig
	Publication             QueueConfig
	IsProducerReturnSuccess bool
	NumberOfRetry           int
}

type Config struct {
	MQ       CommConfig
	GRPC     CommConfig
	Consumer Consumer
	Producer Producer
}

var appConfig Config = Config{
	MQ: CommConfig{
		Domain: localhost,
		Port:   9092,
	},
	GRPC: CommConfig{
		Port:   50002,
		Domain: localhost,
	},
	Consumer: Consumer{
		IsConsumerReturnError: true,
		Confirmation: QueueConfig{
			Topic: constant.ConfirmationNotification,
		},
		Publication: QueueConfig{
			Topic: constant.PublicationNotification,
		},
	},
	Producer: Producer{
		NumberOfRetry:           5,
		IsProducerReturnSuccess: true,
		Confirmation: QueueConfig{
			Topic: constant.ConfirmationEmail,
		},
		Publication: QueueConfig{
			Topic: constant.PublicationEmail,
		},
	},
}

func GetConfig() Config {
	return appConfig
}
