package category

import "catalog/internal/domain/organization"

const TableName = "category"

type Category struct {
	ID            uint   `json:"id" db:"id" gorm:"primary_key"`
	Name          string `json:"name" db:"name"`
	TLeft         uint   `db:"t_left"`
	TRight        uint   `db:"t_right"`
	Organizations []organization.Organization `gorm:"many2many:category2organization"`
}

func (m Category) TableName() string {
	return TableName
}