package category

import (
	"catalog/internal/domain/organization"
	"catalog/internal/pkg/apperror"
	"catalog/internal/pkg/pagination"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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
	BindOrganizations(ctx context.Context, buildingID uint, organizationIDs []uint) (err error)
}

type QueryConditions struct {
	WithOrganizations bool `json:"withOrganization"`
	OrganizationID    uint
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

	if cond.OrganizationID != 0 {
		db = db.Joins("JOIN category2organization cat2org ON cat2org.category_id = category.id").
			Where("cat2org.organization_id = ?", cond.OrganizationID)
	}
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

	if errors.Is(err, apperror.ErrNotFound) && c.ParentID != 0 {
		return 0, errors.Wrap(apperror.ErrNotFound, "parent category not found")
	}

	var parentTRight uint = 0
	if c.ParentID != 0 {
		parentTRight = parentCategory.TRight
		c.TLeft = parentCategory.TRight
		c.TRight = parentCategory.TRight + 1
	} else {
		c.TLeft = 1
		c.TRight = 2
	}

	txErr := m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := m.db.Exec("UPDATE category SET t_left = t_left-2 WHERE id = ?", c.ParentID).Error
		if err != nil {
			return errors.Wrap(err, "cannot update right tree index for category")
		}

		err = m.db.Exec("UPDATE category SET t_left = t_left+2, t_right = t_right+2 WHERE t_right >= ?", parentTRight).Error
		if err != nil {
			return errors.Wrap(err, "cannot update right tree index for category")
		}

		err = tx.Create(c).Error
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
	curCategory, err := m.Get(ctx, c.ID)
	if err != nil {
		return errors.Wrapf(err, "cannot get category id=%d", c.ID)
	}
	newParentCategory, err := m.Get(ctx, c.ParentID)
	if err != nil && err != apperror.ErrNotFound {
		return errors.Wrapf(err, "cannot get category id=%d", newParentCategory.ID)
	}

	var newParentTLeft, newParentTRight uint
	if err == apperror.ErrNotFound {
		count, err := m.Count(ctx, &QueryConditions{})
		if err != nil {
			return err
		}
		maxIndex := count * 2

		newParentTLeft = maxIndex
		newParentTRight = maxIndex + 1
	} else {
		newParentTLeft = newParentCategory.TLeft
		newParentTRight = newParentCategory.TRight
	}

	oldCategory, err := m.Get(ctx, curCategory.ID)
	if err != nil && err != apperror.ErrNotFound {
		return errors.Wrapf(err, "cannot get category id=%d", oldCategory.ID)
	}
	var oldParentTLeft, oldParentTRight uint
	if err == apperror.ErrNotFound {
		count, err := m.Count(ctx, &QueryConditions{})
		if err != nil {
			return err
		}
		maxIndex := count * 2

		oldParentTLeft = maxIndex
		oldParentTRight = maxIndex + 1
	} else {
		oldParentTLeft = oldCategory.TLeft
		oldParentTRight = oldCategory.TRight
	}

	var isMoveToLeft = false
	if oldParentTLeft > newParentTLeft && oldParentTRight > newParentTRight {
		isMoveToLeft = true
	}

	offsetToParentCategory := int64(newParentTRight) - int64(curCategory.TRight) + int64(curCategory.TRight-curCategory.TLeft)
	sizeSubTree := curCategory.TRight - curCategory.TLeft + 1

	return m.db.Transaction(func(tx *gorm.DB) error {
		query := `UPDATE category SET "parent_id"=?,"name"=?,"t_left"=?,"t_right"=? WHERE "id" = ?`
		err = tx.Exec(query, c.ParentID, c.Name, curCategory.TLeft, curCategory.TRight, curCategory.ID).Error
		if err != nil {
			return errors.Wrap(err, "cannot update category subtree")
		}

		query = `UPDATE category SET t_left=t_left+?, t_right=t_right+? WHERE t_left > ? and t_right > ?`
		err = tx.Exec(query, sizeSubTree, sizeSubTree, newParentTLeft, newParentTRight).Error
		if err != nil {
			return errors.Wrap(err, "cannot update category subtree")
		}

		query = `UPDATE category SET t_right=t_right+? WHERE t_left <= ? and t_right >= ?`
		err = tx.Exec(query, sizeSubTree, newParentTLeft, newParentTRight).Error
		if err != nil {
			return errors.Wrap(err, "cannot update category subtree")
		}

		if isMoveToLeft {
			query = `UPDATE category SET t_left=t_left+?,t_right=t_right+? WHERE t_left>=? AND t_right<=?`
			err = tx.Exec(query,
				offsetToParentCategory-int64(sizeSubTree),
				offsetToParentCategory-int64(sizeSubTree),
				curCategory.TLeft+sizeSubTree,
				curCategory.TRight+sizeSubTree,
			).Error
			if err != nil {
				return errors.Wrap(err, "cannot update category subtree")
			}
		} else {
			query = `UPDATE category SET t_left=t_left+?,t_right=t_right+? WHERE t_left>=? AND t_right<=?`
			err = tx.Exec(query,
				offsetToParentCategory,
				offsetToParentCategory,
				curCategory.TLeft,
				curCategory.TRight,
			).Error
			if err != nil {
				return errors.Wrap(err, "cannot update category subtree")
			}
		}

		if isMoveToLeft {
			query = `UPDATE category SET t_left=t_left-?, t_right=t_right-? WHERE t_left > ? and t_right > ?`
			err = tx.Exec(query, sizeSubTree, sizeSubTree, curCategory.TLeft+sizeSubTree, curCategory.TRight+sizeSubTree).Error
			if err != nil {
				return errors.Wrap(err, "cannot update category subtree")
			}
		} else {
			query = `UPDATE category SET t_left=t_left-?, t_right=t_right-? WHERE t_left > ? and t_right > ?`
			err = tx.Exec(query, sizeSubTree, sizeSubTree, curCategory.TLeft, curCategory.TRight).Error
			if err != nil {
				return errors.Wrap(err, "cannot update category subtree")
			}
		}

		if isMoveToLeft {
			query = `UPDATE category SET t_right=t_right-? WHERE t_left < ? AND t_right > ?`
			err = tx.Exec(query, sizeSubTree, curCategory.TLeft+sizeSubTree, curCategory.TRight+sizeSubTree).Error
			if err != nil {
				return errors.Wrap(err, "cannot update category subtree")
			}
		} else {
			query = `UPDATE category SET t_right=t_right-? WHERE t_left < ? AND t_right > ?`
			err = tx.Exec(query, sizeSubTree, curCategory.TLeft, curCategory.TRight).Error
			if err != nil {
				return errors.Wrap(err, "cannot update category subtree")
			}
		}

		return nil
	})
}

func (m repository) Delete(ctx context.Context, id uint) error {
	c, err := m.Get(ctx, id)
	if err != nil {
		return err
	}

	if c.TRight-c.TLeft != 1 {
		return apperror.ErrEntityHasChilds
	}

	return m.db.Debug().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err = tx.Model(c).Association("Organizations").Clear()
		if err != nil {
			return errors.Wrap(err, "cannot delete category")
		}

		err = tx.Delete(c).Error
		if err != nil {
			return errors.Wrap(err, "cannot create category")
		}

		err = m.db.Exec("UPDATE category SET t_left=t_left+2 WHERE id=?", c.ParentID).Error
		if err != nil {
			return errors.Wrap(err, "cannot update category subtree")
		}

		err = m.db.Exec("UPDATE category SET t_left=t_left-2,t_right=t_right-2 WHERE t_right>?", c.TRight).Error
		if err != nil {
			return errors.Wrap(err, "cannot update category subtree")
		}

		return nil
	})
}

func (m repository) Count(ctx context.Context, cond *QueryConditions) (uint, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&Category{}).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot get categorys")
	}

	return uint(count), nil
}

func (m repository) BindOrganizations(ctx context.Context, categoryID uint, organizationIDs []uint) (err error) {
	c, err := m.Get(ctx, categoryID)
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
