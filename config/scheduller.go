package config

import (
	"github.com/go-co-op/gocron"
	"time"
)

var scheduler *gocron.Scheduler

func InitScheduler() {
	var err error
	scheduler = gocron.NewScheduler(time.UTC)
	if err != nil {
		panic(err)
	}
}

func GetScheduler() *gocron.Scheduler {
	return scheduler
}
