package category

import (
	"context"
)

type service struct {
	rep Repository
}

type Service interface {
	Get(ctx context.Context, id uint) (c *Category, err error)
	First(ctx context.Context, cond *Category) (c *Category, err error)
	Query(ctx context.Context, cond *QueryConditions) (organizations []Category, err error)
	Create(ctx context.Context, c *Category) (uint, error)
	Update(ctx context.Context, c *Category) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, cond *QueryConditions) (count uint, err error)
	BindOrganizations(ctx context.Context, buildingID uint, organizationIDs []uint) (err error)
}

func NewService(rep Repository) Service {
	return &service{rep: rep}
}

func (m service) Get(ctx context.Context, id uint) (c *Category, err error) {
	return m.rep.Get(ctx, id)
}

func (m service) First(ctx context.Context, cond *Category) (c *Category, err error) {
	return m.rep.First(ctx, cond)
}

func (m service) Query(ctx context.Context, cond *QueryConditions) (organizations []Category, err error) {
	return m.rep.Query(ctx, cond)
}

func (m service) Create(ctx context.Context, c *Category) (uint, error) {
	return m.rep.Create(ctx, c)
}

func (m service) Update(ctx context.Context, c *Category) error {
	return m.rep.Update(ctx, c)
}

func (m service) Delete(ctx context.Context, id uint) error {
	return m.rep.Delete(ctx, id)
}

func (m service) Count(ctx context.Context, cond *QueryConditions) (count uint, err error) {
	return m.rep.Count(ctx, cond)
}

func (m service) BindOrganizations(ctx context.Context, buildingID uint, organizationIDs []uint) (err error) {
	return m.rep.BindOrganizations(ctx, buildingID, organizationIDs)
}
