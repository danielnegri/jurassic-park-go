CREATE TABLE IF NOT EXISTS cages
(
    id         TEXT                      NOT NULL PRIMARY KEY,
    capacity   INTEGER                   NOT NULL DEFAULT 0,
    active     BOOLEAN,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
);

CREATE INDEX IF NOT EXISTS cages_active_idx ON cages (active);

CREATE TABLE IF NOT EXISTS dinosaurs
(
    id         TEXT                      NOT NULL PRIMARY KEY,
    name       TEXT                      NOT NULL,
    species    TEXT                      NOT NULL,
    cage_id    TEXT                      NOT NULL,

    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    CONSTRAINT dinosaurs_name UNIQUE (name),
    CONSTRAINT dinosaurs_cage_id_fk FOREIGN KEY (cage_id) REFERENCES cages (id)

);

CREATE INDEX IF NOT EXISTS dinosaurs_cage_id_idx ON dinosaurs (cage_id);
CREATE INDEX IF NOT EXISTS dinosaurs_species_idx ON dinosaurs (species);
