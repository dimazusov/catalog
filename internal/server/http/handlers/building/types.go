package building

import (
	"catalog/internal/domain/building"
)

type BuildingsList struct {
	Items []building.Building `json:"items"`
	Count uint                `json:"count"`
}
