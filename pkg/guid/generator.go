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

package guid

import (
	"time"

	"github.com/danielnegri/jurassic-park-go/pkg/guid/base58"
	"github.com/danielnegri/jurassic-park-go/pkg/guid/sonyflake"
)

type (
	// A Generator has sonyflake and encoder.
	Generator struct {
		sf  *sonyflake.Sonyflake
		enc Encoder
	}

	// Settings has setting parameters for indigo.Generator.
	Settings struct {
		StartTime      time.Time
		MachineID      func() (uint16, error)
		CheckMachineID func(uint16) bool
		Encoder        Encoder
	}
)

func DefaultMachineID() (uint16, error) {
	ip, err := getOutboundIP()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

// New settings new a indigo.Generator.
func New(s Settings) *Generator {
	if s.Encoder == nil {
		s.Encoder = base58.StdEncoding
	}
	return &Generator{
		sf: sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime:      s.StartTime,
			MachineID:      s.MachineID,
			CheckMachineID: s.CheckMachineID,
		}),
		enc: s.Encoder,
	}
}

// NextID generates a next unique ID.
func (g *Generator) NextID() (string, error) {
	n, err := g.sf.NextID()
	if err != nil {
		return "", err
	}
	return g.enc.Encode(n), nil
}

func (g *Generator) MustNextID() string {
	return Must(g.NextID())
}

// Must is a helper that wraps a call to a function returning (ID, error)
// and panics if the error is non-nil. It is intended for use in variable
// initializations such as
func Must(id string, err error) string {
	if err != nil {
		panic(err)
	}
	return id
}

// Decompose returns a set of sonyflake ID parts.
func (g *Generator) Decompose(id string) (map[string]uint64, error) {
	b, err := g.enc.Decode(id)
	if err != nil {
		return nil, err
	}
	return sonyflake.Decompose(b), nil
}
