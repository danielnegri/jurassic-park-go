package model

import "time"

const prefixDinosaur = "din"

type Dinosaur struct {
	ID      ID      `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Species Species `json:"species,omitempty"`
	CageID  ID      `json:"cage_id,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func NewDinosaurID(uuid string) ID {
	return NewID(prefixDinosaur, uuid)
}

type DinosaursResource struct {
	Dinosaurs []*Dinosaur `json:"dinosaurs"`
}
