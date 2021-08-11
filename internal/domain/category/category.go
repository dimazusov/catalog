package category

import "catalog/internal/domain/organization"

const RootID = 0
const TableName = "category"

type Category struct {
	ID            uint                        `json:"id" db:"id" gorm:"primary_key"`
	ParentID      uint                        `json:"parentId" db:"parent_id"`
	Name          string                      `json:"name" db:"name"`
	TLeft         uint                        `db:"t_left"`
	TRight        uint                        `db:"t_right"`
	Organizations []organization.Organization `gorm:"many2many:category2organization"`
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
