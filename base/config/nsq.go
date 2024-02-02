package config

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"os"
)

func (cfg Config) InitNsqService() serror.SError {

	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(os.Getenv("SENDER_NSQ_SERVER_REPLACEMENT"), config)

	if err != nil {
		fmt.Println(err.Error())
	}

	cfg.NSQ = producer
	return nil
}
