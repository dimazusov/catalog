package category

import (
	"catalog/internal/domain/category"
)

type CategoryList struct {
	Items []category.Category `json:"items"`
	Count uint                `json:"count"`
}
