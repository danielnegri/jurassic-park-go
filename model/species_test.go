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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpeciesKind(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		species Species
		want    string
	}{
		{name: "empty", species: "", want: KindUnknown},
		{name: "unknown", species: Species("segnosaurus"), want: KindUnknown},
		{name: "Tyrannosaurus", species: Tyrannosaurus, want: KindCarnivore},
		{name: "Velociraptor", species: Velociraptor, want: KindCarnivore},
		{name: "Spinosaurus", species: Spinosaurus, want: KindCarnivore},
		{name: "Megalosaurus", species: Megalosaurus, want: KindCarnivore},
		{name: "Brachiosaurus", species: Brachiosaurus, want: KindHerbivores},
		{name: "Stegosaurus", species: Stegosaurus, want: KindHerbivores},
		{name: "Ankylosaurus", species: Ankylosaurus, want: KindHerbivores},
		{name: "Triceratops", species: Triceratops, want: KindHerbivores},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SpeciesKind(tt.species), "SpeciesKind(%v)", tt.species)
		})
	}
}
