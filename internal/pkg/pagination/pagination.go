package pagination

const DefaultPage = 1
const DefaultPerPage = 10

const MaxPerPage = 1000

type Pagination struct {
	PerPage int `form:"per_page" json:"perPage"`
	Page    int `form:"page" json:"page"`
}

func New(page, perPage int) *Pagination {
	return &Pagination{
		PerPage: perPage,
		Page:    page,
	}
}

func (m Pagination) Validate() error {
	if m.Page < 1 {
		return ErrWrongPage
	}

	if m.PerPage < 1 {
		return ErrWrongCountPerPage
	}

	if m.PerPage > MaxPerPage {
		return ErrExceededMaxCountPerPage
	}

	return nil
}

func (m Pagination) GetLimit() int {
	return m.PerPage
}

func (m Pagination) GetOffset() int {
	return (m.Page - 1) * m.PerPage
}
