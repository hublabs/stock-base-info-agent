package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/hublabs/stock-base-info-agent/factory"
	"github.com/hublabs/stock-base-info-agent/models"
	"github.com/hublabs/stock-base-info-agent/worker"
)

func main() {
	// flag ref:  https://mingrammer.com/gobyexample/command-line-flags/
	flag.Parse()
	factory.Init()

	start := time.Now()
	log.Printf("start: %s \n", start.Format("2006-01-02 15:04:05"))

	// execute etl worker
	execute()

	end := time.Now()
	log.Printf("end: %s \n", end.Format("2006-01-02 15:04:05"))

	duration := end.Sub(start)
	log.Printf("duration: %s \n", duration)
}

func execute() {
	batchJob, err := models.MngBatchJob{}.GetByBatchAndJobName(string(models.Migration), string(models.LocationStore))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	if batchJob.IsUsed {
		mngLocationWorker := worker.MigrationLocationStoreWorker{}.New()
		if err := mngLocationWorker.Run(context.Background()); err != nil {
			log.Println(err.Error())
			panic(err)
		}
	}
}
