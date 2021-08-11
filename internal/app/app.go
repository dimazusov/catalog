package app

import (
	"context"

	"github.com/go-redis/redis/v8"
	minipkg_gorm "github.com/minipkg/db/gorm"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"catalog/internal/cache"
	"catalog/internal/config"
	"catalog/internal/domain/building"
	"catalog/internal/domain/category"
	"catalog/internal/domain/organization"
)

type Domain struct {
	Building     DomainBuilding
	Category     DomainCategory
	Organization DomainOrganization
}

type DomainBuilding struct {
	Repository building.Repository
	Service    building.Service
}

type DomainCategory struct {
	Repository category.Repository
	Service    category.Service
}

type DomainOrganization struct {
	Repository organization.Repository
	Service    organization.Service
}

type App struct {
	cfg    *config.Config
	db     *minipkg_gorm.DB
	redis  *redis.Client
	Cache  cache.Cache
	Domain Domain
}

func New(config *config.Config) *App {
	return &App{cfg: config}
}

func (m *App) DB() *gorm.DB {
	return m.db.DB()
}

func (m *App) LogInfo(data interface{}) error {
	//return m.logger.Info(data)
	return nil
}

func (m *App) Init() error {
	if err := m.initLogger(); err != nil {
		return err
	}
	if err := m.initDB(); err != nil {
		return err
	}
	if err := m.initCache(); err != nil {
		return err
	}
	if err := m.initRepositories(); err != nil {
		return err
	}
	if err := m.initServices(); err != nil {
		return err
	}

	return nil
}

func (m *App) initLogger() (err error) {
	//m.logger, err = logger.New(m.cfg.Logger.Path, m.cfg.Logger.Level)
	return err
}

func (m *App) initDB() (err error) {
	switch m.cfg.Repository.Type {
	case "postgres":
		conn, err := gorm.Open(postgres.Open(m.cfg.DB.Postgres.Dsn), &gorm.Config{})
		if err != nil {
			return errors.Wrapf(err, "cannot connect to postgres")
		}
		m.db = &minipkg_gorm.DB{GormDB: conn}
		break
	}

	return nil
}

func (m *App) initCache() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     m.cfg.Redis.Address,
		Password: m.cfg.Redis.Password,
	})
	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		return errors.New("cannot connect to redis")
	}

	m.Cache = cache.New(rdb)

	return nil
}

func (m *App) initSchemes() (err error) {
	m.db, err = m.db.SchemeInit(&category.Category{})
	m.db, err = m.db.SchemeInit(&building.Building{})
	m.db, err = m.db.SchemeInit(&organization.Organization{})

	return nil
}

func (m *App) initRepositories() (err error) {
	m.Domain.Category.Repository = category.NewRepository(m.db.DB())
	m.Domain.Building.Repository = building.NewRepository(m.db.DB())
	m.Domain.Organization.Repository = organization.NewRepository(m.db.DB())

	return nil
}

func (m *App) initServices() (err error) {
	m.Domain.Category.Service = category.NewService(m.Domain.Category.Repository)
	m.Domain.Building.Service = building.NewService(m.Cache, m.Domain.Building.Repository)
	m.Domain.Organization.Service = organization.NewService(m.Domain.Organization.Repository)

	return nil
}
