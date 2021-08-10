package test_data

import (
	"log"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"catalog/internal/domain/building"
	"catalog/internal/domain/category"
	"catalog/internal/domain/organization"
	"catalog/internal/domain/organization_phone"
	"catalog/internal/pkg/coords"
)

const testCountCategories = 30
const testCountBuildings = 100000
const testCountOrganizations = 100000

const batchSize = 5000

type generator struct {
	db *gorm.DB
}

func NewGenerator(db *gorm.DB) *generator {
	return &generator{
		db: db,
	}
}

func (m generator) GenerateTestData() error {

	return nil
}

func (m generator) GenerateBuildings() {
	log.Println("generate buildings")
	bar := pb.StartNew(testCountOrganizations)
	for i := 0; i < testCountBuildings; i += 5000 {
		buildings := make([]building.Building, 0, batchSize)
		for j := 0; j < batchSize; j++ {
			buildings = append(buildings, building.Building{
				Address: gofakeit.Address().Address,
				Coords: coords.Coords{
					Lat: gofakeit.Address().Latitude,
					Lng: gofakeit.Address().Longitude,
				},
			})
		}
		m.db.Create(&buildings)
	}
	bar.Finish()
}

func (m generator) GenerateOrganizations() error {
	log.Println("generate organizations")

	bar := pb.StartNew(testCountOrganizations)

	buildings := make([]organization.Organization, 0, batchSize)
	for i := 0; i < testCountBuildings; i += 5000 {
		for j := 0; j < batchSize; j++ {
			buildings = append(buildings, organization.Organization{
				Name: gofakeit.Company(),
				Phones: []organization_phone.OrganizationPhone{
					{Number: gofakeit.Phone()},
					{Number: gofakeit.Phone()},
				},
			})
		}
		err := m.db.Create(&buildings)
		if err != nil {
			return errors.New("cannot save buildings")
		}
		buildings = make([]organization.Organization, 0, batchSize)
	}

	bar.Finish()

	return nil
}

func (m generator) GenerateCategories() error {
	log.Println("generate categories")

	ctgs := make([]category.Category, 0, testCountCategories)
	bar := pb.StartNew(testCountCategories)
	for i := 0; i < testCountCategories; i += 5000 {
		ctgs = append(ctgs, category.Category{
			Name: categories[rand.Intn(len(categories))],
		})
	}

	err := m.db.Create(&categories).Error
	if err != nil {
		return err
	}

	bar.Finish()

	return nil
}

// todo: organization to category
// todo: organization to building

// todo: init nested sets