package config

import (
	"fmt"
	"time"

	"github.com/uzzeet/uzzeet-gateway/libs"
	"github.com/uzzeet/uzzeet-gateway/libs/helper"
	"github.com/uzzeet/uzzeet-gateway/libs/helper/serror"
)

func (c *Config) InitTimezone() serror.SError {
	loc := helper.Env(libs.AppTimezone, "Asia/Jakarta")
	local, err := time.LoadLocation(loc)
	if err != nil {
		return serror.NewFromErrorc(err, fmt.Sprintf("failed load location %s", loc))
	}
	time.Local = local
	return nil
}
