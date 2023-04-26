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
