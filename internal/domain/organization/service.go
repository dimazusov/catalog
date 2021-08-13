package organization

import (
	"context"
)

type service struct {
	rep Repository
}

type Service interface {
	Get(ctx context.Context, id uint) (c *Organization, err error)
	First(ctx context.Context, cond *Organization) (o *Organization, err error)
	Query(ctx context.Context, cond *QueryConditions) (organizations []Organization, err error)
	Create(ctx context.Context, o *Organization) (uint, error)
	Update(ctx context.Context, o *Organization) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, cond *QueryConditions) (uint, error)
}

func NewService(rep Repository) Service {
	return &service{rep: rep}
}

func (m service) Get(ctx context.Context, id uint) (o *Organization, err error) {
	return m.rep.Get(ctx, id)
}

func (m service) First(ctx context.Context, cond *Organization) (o *Organization, err error) {
	return m.rep.First(ctx, cond)
}

func (m service) Query(ctx context.Context, cond *QueryConditions) (organizations []Organization, err error) {
	return m.rep.Query(ctx, cond)
}

func (m service) Create(ctx context.Context, o *Organization) (uint, error) {
	return m.rep.Create(ctx, o)
}

func (m service) Update(ctx context.Context, o *Organization) error {
	return m.rep.Update(ctx, o)
}

func (m service) Delete(ctx context.Context, id uint) error {
	return m.rep.Delete(ctx, id)
}

func (m service) Count(ctx context.Context, cond *QueryConditions) (uint, error) {
	return m.rep.Count(ctx, cond)
}
