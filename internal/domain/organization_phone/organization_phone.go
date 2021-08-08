package organization_phone

const TableName = "organization_phone"

type OrganizationPhone struct {
	ID             uint   `json:"id" db:"id" gorm:"primary_key"`
	OrganizationID uint   `json:"organizationId" db:"organization_id"`
	Number         string `json:"number" db:"number"`
}

func (m OrganizationPhone) TableName() string {
	return TableName
}
