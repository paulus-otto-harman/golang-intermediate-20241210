package cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"tugas/service"
)

func CronJob() {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", func() {
		service.ReportOrderExcel()
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.Start()
	fmt.Println("Cron started")

	select {}
}
