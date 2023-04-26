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

package model

import "time"

const (
	MaxCageCapacity = 10

	prefixCage = "cg"
)

type Cage struct {
	ID         ID          `json:"id,omitempty" pg:",pk"`
	Capacity   int         `json:"capacity,omitempty"`
	Allocation int         `json:"allocation,omitempty" pg:"-"`
	Species    Species     `json:"species,omitempty" pg:"-"`
	Dinosaurs  []*Dinosaur `json:"dinosaurs,omitempty" pg:"-"`
	Active     bool        `json:"active,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func NewCageID(uuid string) ID {
	return NewID(prefixCage, uuid)
}

type CagesResource struct {
	Cages []*Cage `json:"cages"`
}
