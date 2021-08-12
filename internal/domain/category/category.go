package category

import "catalog/internal/domain/organization"

const TableName = "category"
const RootID = 0

type Category struct {
	ID            uint                        `json:"id" db:"id" gorm:"primary_key"`
	ParentID      uint                        `json:"parentId" db:"parent_id"`
	Name          string                      `json:"name" db:"name"`
	TLeft         uint                        `json:"-" db:"t_left"`
	TRight        uint                        `json:"-" db:"t_right"`
	Organizations []organization.Organization `json:"organizations,omitempty" gorm:"many2many:category2organization"`
}

func (m Category) TableName() string {
	return TableName
}

func (m *Category) SeTLeftNestedIndex(TLeft uint) {
	m.TLeft = TLeft
}

func (m *Category) SeTRightNestedIndex(TRight uint) {
	m.TRight = TRight
}
