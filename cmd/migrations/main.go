package main

import (
	"catalog/internal/domain/organization"
	"catalog/internal/domain/organization_phone"
	"flag"
	"log"

	"catalog/internal/app"
	"catalog/internal/config"
	"catalog/internal/domain/building"
	"catalog/internal/domain/category"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", `configs/config.yaml`, "Path to configuration file")
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

	db := application.DB()
	err = db.AutoMigrate(building.Building{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(category.Category{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(organization_phone.OrganizationPhone{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(organization.Organization{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("success")
}
