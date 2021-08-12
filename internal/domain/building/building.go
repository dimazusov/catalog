package building

import (
	"catalog/internal/domain/organization"
	"catalog/internal/pkg/coords"
)

const TableName = "building"

type Building struct {
	ID            uint                        `json:"id" db:"id" gorm:"primary_key"`
	Address       string                      `json:"address" db:"address"`
	Coords        coords.Coords               `json:"coords" gorm:"embedded"`
	Organizations []organization.Organization `json:"organizations,omitempty" gorm:"many2many:building2organization"`
}

func (m Building) TableName() string {
	return TableName
}
