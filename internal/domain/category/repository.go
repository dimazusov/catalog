package category

import (
	"context"

	gorm_condition "github.com/minipkg/db/gorm"
	"github.com/minipkg/selection_condition"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"catalog/internal/pkg/apperror"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Get(ctx context.Context, id uint) (c *Category, err error)
	First(ctx context.Context, cond *Category) (c *Category, err error)
	Query(ctx context.Context, cond *selection_condition.SelectionCondition) (organizations []Category, err error)
	Create(ctx context.Context, c *Category) (uint, error)
	Update(ctx context.Context, c *Category) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, cond *selection_condition.SelectionCondition) (count uint, err error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (m repository) Get(ctx context.Context, id uint) (c *Category, err error) {
	err = m.db.WithContext(ctx).First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get by id category")
	}
	return c, nil
}

func (m repository) First(ctx context.Context, cond *Category) (c *Category, err error) {
	err = m.db.WithContext(ctx).Where(cond).First(c).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get category")
	}
	return c, nil
}

func (m repository) Query(ctx context.Context, cond *selection_condition.SelectionCondition) (organizations []Category, err error) {
	db := gorm_condition.Conditions(m.db, cond)

	err = db.WithContext(ctx).Find(&organizations).Error
	if err != nil {
		return nil, errors.Wrap(err, "cannot get categorys")
	}
	return organizations, nil
}

func (m repository) Create(ctx context.Context, c *Category) (uint, error) {
	err := m.db.WithContext(ctx).Create(c).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot create category")
	}
	return c.ID, nil
}

func (m repository) Update(ctx context.Context, c *Category) error {
	err := m.db.WithContext(ctx).Create(c).Error
	if err != nil {
		return errors.Wrap(err, "cannot update category")
	}
	return nil
}

func (m repository) Delete(ctx context.Context, id uint) error {
	err := m.db.WithContext(ctx).Delete(&Category{ID: id}).Error
	if err != nil {
		return errors.Wrap(err, "cannot delete category")
	}
	return nil
}

func (m repository) Count(ctx context.Context, cond *selection_condition.SelectionCondition) (uint, error) {
	db := gorm_condition.Conditions(m.db, cond)

	var count int64
	err := db.WithContext(ctx).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot get categorys")
	}

	return uint(count), nil
}
