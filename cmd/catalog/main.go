package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"catalog/internal/app"
	"catalog/internal/config"
	http_server "catalog/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", config.DefaultConfigPath, "Path to configuration file")
	flag.Parse()
}

func main() {
	cfg, err := config.New(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	application := app.New(cfg)

	err = application.Init()
	if err != nil {
		log.Fatalln(err)
	}

	srv := http_server.NewServer(cfg, application)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP)

		<-signals
		signal.Stop(signals)

		err := srv.Stop(context.Background())
		if err != nil {
			log.Println(err)
		}
	}()

	err = srv.Start(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
}
