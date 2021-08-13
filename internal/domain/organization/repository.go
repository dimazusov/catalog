package organization

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"catalog/internal/pkg/apperror"
	"catalog/internal/pkg/pagination"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Get(ctx context.Context, id uint) (c *Organization, err error)
	First(ctx context.Context, cond *Organization) (c *Organization, err error)
	Query(ctx context.Context, cond *QueryConditions) (organizations []Organization, err error)
	Create(ctx context.Context, c *Organization) (uint, error)
	Update(ctx context.Context, c *Organization) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, cond *QueryConditions) (uint, error)
}

type QueryConditions struct {
	WithPhones bool `json:"withPhones"`
	*pagination.Pagination
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (m repository) Get(ctx context.Context, id uint) (o *Organization, err error) {
	o = &Organization{}
	err = m.db.WithContext(ctx).First(o, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get by id organization")
	}
	return o, nil
}

func (m repository) First(ctx context.Context, cond *Organization) (o *Organization, err error) {
	o = &Organization{}
	err = m.db.WithContext(ctx).Where(cond).First(o).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get organization")
	}
	return o, nil
}

func (m repository) Query(ctx context.Context, cond *QueryConditions) (organizations []Organization, err error) {
	db := m.db.WithContext(ctx).
		Offset(cond.Pagination.GetOffset()).
		Limit(cond.Pagination.GetLimit())

	if cond.WithPhones {
		db = db.Preload("Phones")
	}

	err = db.WithContext(ctx).Find(&organizations).Error
	if err != nil {
		return nil, errors.Wrap(err, "cannot get organizations")
	}
	return organizations, nil
}

func (m repository) Create(ctx context.Context, o *Organization) (uint, error) {
	err := m.db.WithContext(ctx).Create(o).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot create organization")
	}
	return o.ID, nil
}

func (m repository) Update(ctx context.Context, o *Organization) error {
	err := m.db.WithContext(ctx).Save(o).Error
	if err != nil {
		return errors.Wrap(err, "cannot update organization")
	}
	return nil
}

func (m repository) Delete(ctx context.Context, id uint) error {
	org := Organization{ID: id}
	err := m.db.WithContext(ctx).Model(&org).Select(clause.Associations).Delete(&org).Error
	if err != nil {
		return errors.Wrap(err, "cannot delete organization")
	}
	return nil
}

func (m repository) Count(ctx context.Context, cond *QueryConditions) (uint, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&Organization{}).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot get organizations count")
	}

	return uint(count), nil
}
