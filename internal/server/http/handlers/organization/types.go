package organization

import (
	"catalog/internal/domain/organization"
)

type OrganizationList struct {
	Items []organization.Organization `json:"items"`
	Count uint                        `json:"count"`
}
