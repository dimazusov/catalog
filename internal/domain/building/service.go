package building

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"catalog/internal/cache"
	"catalog/internal/pkg/apperror"
	"catalog/internal/pkg/hasher"
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
	BindOrganizations(ctx context.Context, buildingID uint, organizationIDs []uint) (err error)
}

func NewService(cache cache.Cache, rep Repository) Service {
	return &service{rep: rep, cache: cache}
}

func (m service) Get(ctx context.Context, id uint) (b *Building, err error) {
	b, err = m.getBuildingFromCache(id)
	if err == nil {
		return b, nil
	}
	if !errors.Is(err, apperror.ErrNotFound) {
		return nil, err
	}

	if b, err = m.rep.Get(ctx, id); err != nil {
		return nil, err
	}

	if err = m.addBuildingToCache(b); err != nil {
		return nil, err
	}

	return b, nil
}

func (m service) First(ctx context.Context, cond *Building) (b *Building, err error) {
	if b, err = m.rep.Get(ctx, cond.ID); err != nil {
		return nil, err
	}

	return b, nil
}

func (m service) Query(ctx context.Context, cond *QueryConditions) (buildings []Building, err error) {
	buildings, err = m.getBuildingsFromCache(cond)
	if err == nil {
		return buildings, nil
	}
	if !errors.Is(err, apperror.ErrNotFound) {
		return nil, err
	}

	if buildings, err = m.rep.Query(ctx, cond); err != nil {
		return nil, err
	}

	err = m.addBuildingsToCache(cond, buildings)
	if err != nil {
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

func (m service) BindOrganizations(ctx context.Context, buildingID uint, organizationIDs []uint) (err error) {
	return m.rep.BindOrganizations(ctx, buildingID, organizationIDs)
}

func (m service) getBuildingFromCache(buildingId uint) (*Building, error) {
	cacheKey, err := hasher.New().GetHashFromStruct(Building{ID: buildingId})
	if err != nil {
		return nil, errors.Wrap(err, "cannot get cache key")
	}

	val, err := m.cache.Get(cacheKey)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get building from cache")
	}

	b := &Building{}
	err = b.UnmarshalJSON([]byte(val.(string)))
	if err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal building from cache")
	}

	return b, nil
}

func (m service) addBuildingToCache(b *Building) error {
	cacheKey, err := hasher.New().GetHashFromStruct(Building{ID: b.ID})
	if err != nil {
		return errors.Wrap(err, "cannot get cache key")
	}

	data, err := b.MarshalJSON()
	if err != nil {
		return errors.Wrapf(err, "cannot marshal building %#v", *b)
	}

	err = m.cache.Set(cacheKey, data)
	if err != nil {
		return errors.Wrapf(err, "cannot set building to cache %#v", data)
	}

	return nil
}

func (m service) getBuildingsFromCache(cond *QueryConditions) ([]Building, error) {
	cacheKey, err := hasher.New().GetHashFromStruct(*cond)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get cache key")
	}

	val, err := m.cache.Get(cacheKey)
	if err != nil {
		return nil, err
	}

	buildings := []Building{}

	err = json.Unmarshal([]byte(val.(string)), &buildings)
	if err != nil {
		return nil, err
	}

	return buildings, nil
}

func (m service) addBuildingsToCache(cond *QueryConditions, buildings []Building) error {
	data, err := json.Marshal(&buildings)
	if err != nil {
		return err
	}

	cacheKey, err := hasher.New().GetHashFromStruct(*cond)
	if err != nil {
		return errors.Wrap(err, "cannot get cache key")
	}

	err = m.cache.Set(cacheKey, string(data))
	if err != nil {
		return err
	}

	return nil
}
