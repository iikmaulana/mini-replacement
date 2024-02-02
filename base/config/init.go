package config

import (
	"fmt"
	"github.com/nsqio/go-nsq"

	"github.com/jmoiron/sqlx"
	"github.com/uzzeet/uzzeet-gateway/controller"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
	"github.com/uzzeet/uzzeet-gateway/service"
)

type AppConfig interface {
	Init() Config
}

type Config struct {
	Version  string
	DB       *sqlx.DB
	Registry *controller.Registry
	Server   *service.Server
	Gateway  *service.Service
	CK       *sqlx.DB
	NSQ      *nsq.Producer
}

func Init() Config {
	var cfg Config

	errx := cfg.initGateway()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitTimezone()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitCockroachdb()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitNsqService()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitService()
	if errx != nil {
		errx.Panic()
	}

	fmt.Println("Server is running ..")
	return cfg
}

func New() Config {
	return Config{
		Version: "v1.1.5",
	}
}

func (ox *Config) Start() (errx serror.SError) {
	ch := make(chan bool)
	go func() {
		err := ox.Server.Start()
		if err != nil {
			errx = serror.NewFromErrorc(err, "Cannot starting server")
		}

		ch <- false
	}()

	<-ch
	return errx
}
