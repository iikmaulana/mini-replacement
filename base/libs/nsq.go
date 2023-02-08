package lib

import (
	"github.com/nsqio/go-nsq"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"os"
)

const RealmIdDealer = "3bdb348a-9b79-4738-b804-8b872bde8e7e"
const RealmIdCustomer = "9ceb58f2-66ed-49e5-93c7-0b897c1ea8c7"

type PayloadNsq struct {
	RequestID    string `json:"request_id"`
	Time         string `json:"time"`
	Service      string `json:"service"`
	DatabaseName string `json:"database_name"`
	Payload      string `json:"payload"`
	Query        string `json:"query"`
}

func SendNSQUsecase(form []byte) (serr serror.SError) {

	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(os.Getenv("SENDER_NSQ_SERVER_REPLACEMENT"), config)
	if err != nil {
		return serror.NewFromError(err)
	}
	defer producer.Stop()

	topic := os.Getenv("SENDER_NSQ_TOPIC_REPLACEMENT")
	err = producer.Publish(topic, form)
	if err != nil {
		return serror.NewFromError(err)
	}

	return nil
}
