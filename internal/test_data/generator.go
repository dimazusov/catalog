package test_data

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/cheggaaa/pb/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

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
	//db.Logger = newLogger()
	return &generator{
		db: db,
	}
}

func newLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
}

func (m generator) GenerateTestData() (err error) {
	if err := m.generateBuildings(); err != nil {
		return err
	}

	if err := m.generateCategories(); err != nil {
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
	bar := pb.StartNew(len(categories))

	for _, catName := range categories {
		c := category.Category{Name: catName}
		err := m.db.Create(&c).Error
		if err != nil {
			return errors.Wrap(err, "cannot create categories")
		}

		childsCategories := make([]category.Category, 0, maxCountSubCategories)
		for i := 0; i < maxCountSubCategories; i += 1 {
			childsCategories = append(childsCategories, category.Category{
				ParentID: c.ID,
				Name:     categories[rand.Intn(len(categories))],
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

//func (m generator) GenerateOrganizations() error {
//	bar := pb.StartNew(testCountOrganizations)
//
//	buildings := make([]organization.Organization, 0, batchSize)
//	for i := 0; i < testCountBuildings; i += 5000 {
//		for j := 0; j < batchSize; j++ {
//			buildings = append(buildings, organization.Organization{
//				Name: gofakeit.Company(),
//				Phones: []organization_phone.OrganizationPhone{
//					{Number: gofakeit.Phone()},
//					{Number: gofakeit.Phone()},
//				},
//			})
//		}
//		err := m.db.Create(&buildings).Error
//		if err != nil {
//			return errors.Wrap(err, "cannot save organization")
//		}
//		buildings = make([]organization.Organization, 0, batchSize)
//
//		bar.Add(len(buildings))
//	}
//
//	bar.Finish()
//
//	return nil
//}
//
//
//func (m generator) GenerateCategories2Organizations() (err error) {
//	//ctgs := []category.Category{}
//	//err = m.db.Where("parent_id != ? ", category.RootID).Find(&ctgs).Error
//	//if err != nil {
//	//	return err
//	//}
//	//
//	//var countOrganizations int64
//	//err = m.db.Model(&organization.Organization{}).Count(&countOrganizations).Error
//	//if err != nil {
//	//	return err
//	//}
//	//
//	//for _, c := range ctgs {
//	//	categories[rand.Intn(countOrganizations)]
//	//	c.ID
//	//
//	//}
//
//	log.Println(m.getRandomID(123))
//
//	c := category.Category{
//		ParentID: 0,
//		Name: "test",
//		Organizations: []organization.Organization{
//			{
//				Name: "organization name", Phones: []organization_phone.OrganizationPhone{
//					{Number: "41258-124-123"},
//				},
//			},
//		},
//	}
//
//	err = m.db.Create(&c).Error
//	if err != nil {
//		log.Println(err)
//	}
//
//	log.Println(c.ID)
//
//
//	// get count organization
//	// categories without childs
//
//	// for categories {
//	// rand organization given organizationIDs
//	// }
//	// save
//
//	return nil
//}
//
//func (m generator) GenerateBuildings2Organizations() (err error) {
//	// get count organization
//	// get count buildings
//
//	// for categories {
//	// rand organization given organizationIDs
//	// }
//	// save
//
//	return nil
//}
//
//func (m generator) getRandomID(total uint) int {
//	val := rand.Intn(int(total))
//	if val == 0 {
//		return 1
//	}
//	return val
//}
//// todo: organization to category
//// todo: organization to building
//
//// todo: init nested sets
