package config

type Config struct {
	KafkaHost string
	KafkaPort string

	KafkaTopic string
}

func NewConfig(kafkaHost, kafkaPort, kafkaTopic string) *Config {
	return &Config{
		KafkaHost:  kafkaHost,
		KafkaPort:  kafkaPort,
		KafkaTopic: kafkaTopic,
	}
}
