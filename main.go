package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bugsnag/bugsnag-go"
	"github.com/seannguyen/coin-tracker/internal/jobs"
	"github.com/seannguyen/coin-tracker/internal/utilities"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	initConfigs()

	stopSignal := getStopSignalChan()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	startJob(stopSignal, wg)

	wg.Wait()
	log.Println("Server shuted down")
}

func initConfigs() {
	log.SetOutput(os.Stdout)

	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}

	if utilities.IsDevelopment() {
		boil.DebugMode = true
	}

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       viper.GetString("BUGSNAG_API_KEY"),
		ReleaseStage: viper.GetString("ENV"),
	})

	log.Println("Initialized configs")
}

func getStopSignalChan() <-chan struct{} {
	terminate := make(chan os.Signal)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM)
	stop := make(chan struct{})

	go func() {
		<-terminate
		close(stop)
	}()

	return stop
}

func startJob(stop <-chan struct{}, wg *sync.WaitGroup) {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		execJob()
		for {
			select {
			case <-stop:
				ticker.Stop()
				wg.Done()
				log.Println("Interval job stopped")
				return
			case <-ticker.C:
				execJob()
			}
		}
	}()
	log.Println("Interval job started")
}

func execJob() {
	defer func() {
		if r := recover(); r != nil {
			bugsnag.Notify(r.(error))
		}
	}()

	jobs.SnapshotBalances()
}
