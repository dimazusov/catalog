package test_data

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"math/rand"
	"os"

	"catalog/internal/domain/building"
	"catalog/internal/domain/category"
	"catalog/internal/domain/organization"
	"catalog/internal/domain/organization_phone"
	"catalog/internal/pkg/coords"
)

const maxCountSubCategories = 20
const maxCategory2Organization = 20
const maxBuilding2Organization = 20
const countBuildings = 100000

const batchSize = 5000

type generator struct {
	db *gorm.DB
}

func NewGenerator(db *gorm.DB) *generator {
	db.Logger = newLogger()
	return &generator{
		db: db,
	}
}

func newLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
}

func (m generator) GenerateTestData() (err error) {
	log.Println("generate buildings")
	if err := m.generateBuildings(); err != nil {
		return err
	}

	log.Println("generate categories")
	if err := m.generateCategories(); err != nil {
		return err
	}

	log.Println("init nested sets")
	if _, err := m.initCategoryNestedSets(category.RootID, 0); err != nil {
		return err
	}

	return nil
}

func (m generator) generateBuildings() error {
	buildings := make([]building.Building, 0, batchSize)
	bar := pb.StartNew(countBuildings)
	for i := 0; i < countBuildings; i += 5000 {
		for j := 0; j < batchSize; j++ {
			buildings = append(buildings, building.Building{
				Address: gofakeit.Address().Address,
				Coords: coords.Coords{
					Lat: gofakeit.Address().Latitude,
					Lng: gofakeit.Address().Longitude,
				},
				Organizations: m.generateOrganizations(rand.Intn(maxBuilding2Organization)),
			})

			err := m.db.Create(&buildings).Error
			if err != nil {
				return err
			}

			bar.Add(len(buildings))

			buildings = make([]building.Building, 0, batchSize)
		}
	}
	bar.Finish()

	return nil
}

func (m generator) generateCategories() (err error) {
	bar := pb.StartNew(len(categoryNames))

	for _, ctgName := range categoryNames {
		c := category.Category{Name: ctgName}
		err := m.db.Create(&c).Error
		if err != nil {
			return errors.Wrap(err, "cannot create categories")
		}

		childsCategories := make([]category.Category, 0, maxCountSubCategories)
		for i := 0; i < maxCountSubCategories; i += 1 {
			childsCategories = append(childsCategories, category.Category{
				ParentID:      c.ID,
				Name:          categoryNames[rand.Intn(len(categoryNames))],
				Organizations: m.generateOrganizations(rand.Intn(maxCategory2Organization)),
			})
		}

		err = m.db.Create(&childsCategories).Error
		if err != nil {
			return errors.Wrap(err, "cannot create childs categories")
		}
		bar.Add(1)
	}

	bar.Finish()

	return nil
}

func (m generator) generateOrganizations(count int) []organization.Organization {
	organizations := make([]organization.Organization, 0, count)

	for i := 0; i < count; i++ {
		organizations = append(organizations, organization.Organization{
			Name: gofakeit.Company(),
			Phones: []organization_phone.OrganizationPhone{
				{Number: gofakeit.Phone()},
				{Number: gofakeit.Phone()},
			},
		})
	}

	return organizations
}

func (m generator) initCategoryNestedSets(parentCategoryID, beginNestedIndex uint) (currentNestedIndex uint, err error) {
	categories := []category.Category{}

	err = m.db.Where("parent_id = ?", parentCategoryID).Find(&categories).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot get categories")
	}

	currentNestedIndex = beginNestedIndex

	for i, _ := range categories {
		currentNestedIndex += 1
		categories[i].SeTLeftNestedIndex(currentNestedIndex)

		currentNestedIndex, err = m.initCategoryNestedSets(categories[i].ID, currentNestedIndex)
		if err != nil {
			return 0, err
		}

		currentNestedIndex += 1
		categories[i].SeTRightNestedIndex(currentNestedIndex)

		err = m.db.Save(&categories[i]).Error
		if err != nil {
			return 0, errors.Wrapf(err, "cannot save category, id: %d", categories[i].ID)
		}
	}

	return currentNestedIndex, nil
}
