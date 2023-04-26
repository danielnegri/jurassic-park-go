-- Copyright 2023 The Jurassic Park Authors
--
-- Licensed under the AGPL, Version 3.0 (the "License");
-- you may not use this file except in compliance with the License.
-- You may obtain a copy of the License at
--
--     https://www.gnu.org/licenses/agpl-3.0.en.html
--
-- Unless required by applicable law or agreed to in writing, software
-- distributed under the License is distributed on an "AS IS" BASIS,
-- WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-- See the License for the specific language governing permissions and
-- limitations under the License.

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
