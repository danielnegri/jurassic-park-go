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
