package storage

import (
	"context"

	"github.com/danielnegri/jurassic-park-go/model"
)

type Storage interface {
	Close() error
	Check(ctx context.Context) error

	CreateCage(ctx context.Context, cage *model.Cage) error

	UpdateCage(ctx context.Context, id model.ID, updater CageUpdater) error

	GetCage(ctx context.Context, id model.ID) (*model.Cage, error)

	ListCages(ctx context.Context, params ListCageParams) ([]*model.Cage, error)
}

type (
	CageUpdater func(old *model.Cage) (*model.Cage, error)

	ListCageParams struct {
		Pagination *Pagination
		Status     string
	}
)

const (
	PaginationLimit = 100
)

// Pagination is passed as a parameter to limit the total of rows.
type Pagination struct {
	Limit  int
	Offset int
}

func NewPagination(perPage, page int) *Pagination {
	if perPage >= PaginationLimit {
		perPage = PaginationLimit
	}

	return &Pagination{
		Limit:  perPage,
		Offset: page * perPage,
	}
}
