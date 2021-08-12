package category

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"catalog/internal/pkg/apperror"
	"catalog/internal/pkg/pagination"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Get(ctx context.Context, id uint) (c *Category, err error)
	First(ctx context.Context, cond *Category) (c *Category, err error)
	Query(ctx context.Context, cond *QueryConditions) (organizations []Category, err error)
	Create(ctx context.Context, c *Category) (uint, error)
	Update(ctx context.Context, c *Category) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, cond *QueryConditions) (count uint, err error)
}

type QueryConditions struct {
	WithOrganizations bool `json:"withOrganization"`
	*pagination.Pagination
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (m repository) Get(ctx context.Context, id uint) (c *Category, err error) {
	c = &Category{}
	err = m.db.WithContext(ctx).Preload("Organizations").First(c, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get by id category")
	}
	return c, nil
}

func (m repository) First(ctx context.Context, cond *Category) (c *Category, err error) {
	c = &Category{}
	err = m.db.WithContext(ctx).Preload("Organizations").Where(cond).First(c).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get category")
	}
	return c, nil
}

func (m repository) Query(ctx context.Context, cond *QueryConditions) (organizations []Category, err error) {
	db := m.db.WithContext(ctx).
		Offset(cond.Pagination.GetOffset()).
		Limit(cond.Pagination.GetLimit())

	if cond.WithOrganizations {
		db = db.Preload("Organizations")
	}

	err = db.Find(&organizations).Error
	if err != nil {
		return nil, errors.Wrap(err, "cannot get categorys")
	}
	return organizations, nil
}

func (m repository) Create(ctx context.Context, c *Category) (newID uint, err error) {
	parentCategory, err := m.Get(ctx, c.ParentID)
	if err != nil && err != apperror.ErrNotFound {
		return 0, err
	}

	txErr := m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err = m.db.Model(c).
			Where("t_right >= ?", parentCategory.TRight).
			Update("t_right", "t_right+2").Error
		if err != nil {
			return errors.Wrap(err, "cannot update right tree index for category")
		}

		err := tx.Create(c).Error
		if err != nil {
			return errors.Wrap(err, "cannot create category")
		}
		return nil
	})

	if txErr != nil {
		return 0, errors.Wrap(err, "cannot create category")
	}

	return c.ID, nil
}

func (m repository) Update(ctx context.Context, c *Category) error {
	err := m.db.WithContext(ctx).Save(c).Error
	if err != nil {
		return errors.Wrap(err, "cannot update category")
	}
	return nil
}

func (m repository) Delete(ctx context.Context, id uint) error {
	c, err := m.Get(ctx, id)
	if err != nil {
		return err
	}

	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := m.db.Model(c).
			Where("t_right > ?", c.TRight).
			Update("t_right", "t_right-2").Error
		if err != nil {
			return errors.Wrap(err, "cannot update right tree index for category")
		}

		err = tx.Where("t_left >= ?", c.TLeft).
			Where("t_right <= ?", c.TRight).
			Delete(c).Error
		if err != nil {
			return errors.Wrap(err, "cannot create category")
		}
		return nil
	})
}

func (m repository) Count(ctx context.Context, cond *QueryConditions) (uint, error) {
	var count int64
	err := m.db.WithContext(ctx).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot get categorys")
	}

	return uint(count), nil
}
