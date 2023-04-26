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
