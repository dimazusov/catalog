package building

import (
	"context"

	"github.com/minipkg/selection_condition"
)

type service struct {
	rep Repository
}

type Service interface {
	Get(ctx context.Context, id uint) (b *Building, err error)
	First(ctx context.Context, cond *Building) (b *Building, err error)
	Query(ctx context.Context, cond *selection_condition.SelectionCondition) (buildings []Building, err error)
	Create(ctx context.Context, b *Building) (uint, error)
	Update(ctx context.Context, b *Building) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, cond *selection_condition.SelectionCondition) (uint, error)
}

func NewService(rep Repository) Service {
	return &service{rep: rep}
}

func (m service) Get(ctx context.Context, id uint) (b *Building, err error) {
	return m.rep.Get(ctx, id)
}

func (m service) First(ctx context.Context, cond *Building) (b *Building, err error) {
	return m.rep.First(ctx, cond)
}

func (m service) Query(ctx context.Context, cond *selection_condition.SelectionCondition) (buildings []Building, err error) {
	return m.rep.Query(ctx, cond)
}

func (m service) Create(ctx context.Context, b *Building) (uint, error) {
	return m.rep.Create(ctx, b)
}

func (m service) Update(ctx context.Context, b *Building) error {
	return m.rep.Update(ctx, b)
}

func (m service) Delete(ctx context.Context, id uint) error {
	return m.rep.Delete(ctx, id)
}

func (m service) Count(ctx context.Context, cond *selection_condition.SelectionCondition) (uint, error) {
	return m.rep.Count(ctx, cond)
}