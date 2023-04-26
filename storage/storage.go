// Copyright 2023 The Jurassic Park Authors
//
// Licensed under the AGPL, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.gnu.org/licenses/agpl-3.0.en.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
