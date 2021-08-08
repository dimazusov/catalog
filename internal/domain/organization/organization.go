package organization

import "catalog/internal/domain/organization_phone"

const TableName = "organization"

type Organization struct {
	ID     uint                                   `json:"id" db:"id" gorm:"primary_key"`
	Name   string                                 `json:"name" db:"name"`
	Phones []organization_phone.OrganizationPhone `json:"phones" gorm:"foreignKey:OrganizationID"`
}

func (m Organization) TableName() string {
	return TableName
}
