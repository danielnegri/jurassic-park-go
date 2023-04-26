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

type Species string

const (
	KindCarnivore  = "carnivore"
	KindHerbivores = "herbivores"
	KindUnknown    = "unknown"
)

const (
	Tyrannosaurus Species = "tyrannosaurus"
	Velociraptor  Species = "velociraptor"
	Spinosaurus   Species = "spinosaurus"
	Megalosaurus  Species = "megalosaurus"
)

const (
	Brachiosaurus Species = "brachiosaurus"
	Stegosaurus   Species = "stegosaurus"
	Ankylosaurus  Species = "ankylosaurus"
	Triceratops   Species = "triceratops"
)

var kinds = map[Species]string{
	Tyrannosaurus: KindCarnivore,
	Velociraptor:  KindCarnivore,
	Spinosaurus:   KindCarnivore,
	Megalosaurus:  KindCarnivore,
	Brachiosaurus: KindHerbivores,
	Stegosaurus:   KindHerbivores,
	Ankylosaurus:  KindHerbivores,
	Triceratops:   KindHerbivores,
}

func SpeciesKind(s Species) string {
	if kind, exists := kinds[s]; exists {
		return kind
	}

	return KindUnknown
}
