package building

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"catalog/internal/domain/organization"
	"catalog/internal/pkg/apperror"
	"catalog/internal/pkg/pagination"
)

type Repository interface {
	Get(ctx context.Context, id uint) (c *Building, err error)
	First(ctx context.Context, cond *Building) (c *Building, err error)
	Query(ctx context.Context, cond *QueryConditions) (buildings []Building, err error)
	Create(ctx context.Context, c *Building) (uint, error)
	Update(ctx context.Context, c *Building) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, cond *QueryConditions) (uint, error)
	BindOrganizations(ctx context.Context, buildingID uint, organizationIDs []uint) (err error)
}

type QueryConditions struct {
	WithOrganizations bool `form:"with_organization" json:"withOrganization"`
	OrganizationID    uint
	*pagination.Pagination
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (m repository) Get(ctx context.Context, id uint) (b *Building, err error) {
	b = &Building{}
	err = m.db.WithContext(ctx).First(b, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get by id building")
	}
	return b, nil
}

func (m repository) First(ctx context.Context, cond *Building) (b *Building, err error) {
	b = &Building{}
	err = m.db.WithContext(ctx).Where(cond).First(b).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get building")
	}
	return b, nil
}

func (m repository) Query(ctx context.Context, cond *QueryConditions) (buildings []Building, err error) {
	db := m.db.Debug().WithContext(ctx).
		Offset(cond.Pagination.GetOffset()).
		Limit(cond.Pagination.GetLimit())

	if cond.OrganizationID != 0 {
		db = db.Joins("JOIN building2organization bld2org ON bld2org.building_id = building.id").
			Where("bld2org.organization_id = ?", cond.OrganizationID)
	}
	if cond.WithOrganizations {
		db = db.Preload("Organizations")
	}

	err = db.Find(&buildings).Error
	if err != nil {
		return nil, errors.Wrap(err, "cannot get buildings")
	}
	return buildings, nil
}

func (m repository) Create(ctx context.Context, c *Building) (uint, error) {
	err := m.db.WithContext(ctx).Create(c).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot create building")
	}
	return c.ID, nil
}

func (m repository) Update(ctx context.Context, c *Building) error {
	err := m.db.WithContext(ctx).Save(c).Error
	if err != nil {
		return errors.Wrap(err, "cannot update building")
	}
	return nil
}

func (m repository) Delete(ctx context.Context, id uint) error {
	b, err := m.Get(ctx, id)
	if err != nil {
		return err
	}

	err = m.db.WithContext(ctx).Model(b).Select(clause.Associations).Delete(b).Error
	if err != nil {
		return errors.Wrap(err, "cannot delete building")
	}
	return nil
}

func (m repository) Count(ctx context.Context, cond *QueryConditions) (uint, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&Building{}).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot get count buildings")
	}

	return uint(count), nil
}

func (m repository) BindOrganizations(ctx context.Context, buildingID uint, organizationIDs []uint) (err error) {
	c, err := m.Get(ctx, buildingID)
	if err != nil {
		return err
	}

	organizations := make([]organization.Organization, 0, len(organizationIDs))
	for _, orgID := range organizationIDs {
		organizations = append(organizations, organization.Organization{ID: orgID})
	}

	err = m.db.Model(c).Association("Organizations").Replace(organizations)
	if err != nil {
		return err
	}

	return nil
}
