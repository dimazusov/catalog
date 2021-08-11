package main

import (
	"catalog/internal/app"
	"catalog/internal/config"
	"catalog/internal/test_data"
	"flag"
	"log"
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

	generator := test_data.NewGenerator(application.DB())
	err = generator.GenerateTestData()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("generate successful")
}