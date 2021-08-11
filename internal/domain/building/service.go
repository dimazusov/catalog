package building

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/pkg/errors"

	"catalog/internal/pkg/apperror"
	"catalog/internal/cache"
)

type service struct {
	rep   Repository
	cache cache.Cache
}

type Service interface {
	Get(ctx context.Context, id uint) (b *Building, err error)
	First(ctx context.Context, cond *Building) (b *Building, err error)
	Query(ctx context.Context, cond *QueryConditions) (buildings []Building, err error)
	Create(ctx context.Context, b *Building) (uint, error)
	Update(ctx context.Context, b *Building) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, cond *QueryConditions) (uint, error)
}

func NewService(cache cache.Cache, rep Repository) Service {
	return &service{rep: rep, cache: cache}
}

func (m service) Get(ctx context.Context, id uint) (b *Building, err error) {
	key := m.getBuildingCacheKey(id)

	val, err := m.cache.Get(key)
	if err == nil {
		bld := val.(Building)
		return &bld, nil
	}

	if !errors.Is(err, apperror.ErrNotFound) {
		return nil, err
	}

	if b, err = m.rep.Get(ctx, id); err != nil {
		return nil, err
	}

	if err = m.cache.Set(key, *b); err != nil {
		return nil, err
	}

	return b, nil
}

func (m service) First(ctx context.Context, cond *Building) (b *Building, err error) {
	key := m.getBuildingCacheKey(cond.ID)

	val, err := m.cache.Get(key)
	if err == nil {
		bld := val.(Building)
		return &bld, nil
	}

	if !errors.Is(err, apperror.ErrNotFound) {
		return nil, err
	}

	if b, err = m.rep.Get(ctx, cond.ID); err != nil {
		return nil, err
	}

	if err = m.cache.Set(key, *b); err != nil {
		return nil, err
	}

	return b, nil
}

func (m service) Query(ctx context.Context, cond *QueryConditions) (buildings []Building, err error) {
	key := m.getBuildingsCacheKey(cond)

	val, err := m.cache.Get(key)
	buildings = val.([]Building)
	if err == nil {
		return buildings, nil
	}

	if !errors.Is(err, apperror.ErrNotFound) {
		return nil, err
	}

	if buildings, err = m.rep.Query(ctx, cond); err != nil {
		return nil, err
	}

	if err = m.cache.Set(key, buildings); err != nil {
		return nil, err
	}

	return buildings, nil
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

func (m service) Count(ctx context.Context, cond *QueryConditions) (uint, error) {
	return m.rep.Count(ctx, cond)
}

func (m service) getBuildingCacheKey(buildingId uint) string {
	return fmt.Sprintf("%s.%d", TableName, buildingId)
}

func (m service) getBuildingsCacheKey(cond *QueryConditions) string {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(cond)
	return string(b.Bytes())
}