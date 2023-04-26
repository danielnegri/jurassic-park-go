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

package postgres

import (
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/danielnegri/jurassic-park-go/pkg/app"
	"github.com/danielnegri/jurassic-park-go/pkg/guid"
	"github.com/danielnegri/jurassic-park-go/pkg/log"
	"github.com/go-pg/pg/v10"
)

var (
	postgres *Postgres
	uuid     *guid.Generator
)

func getDatabaseURL() (string, bool) {
	return os.LookupEnv("TEST_DATABASE_URL")
}

func shouldSkip() bool {
	_, exists := getDatabaseURL()
	return !exists
}

func setup(t *testing.T) {
	t.Helper()

	gofakeit.Seed(0)

	log.SetLogger(log.New("debug", "json"))

	if postgres == nil {
		pg, err := newTestPostgres(t)
		if err != nil {
			t.Fatal(err)
		}

		postgres = pg
	}

	if uuid == nil {
		uuid = guid.New(guid.Settings{StartTime: app.StartDate()})
	}
}

func newTestPostgres(t *testing.T) (*Postgres, error) {
	databaseURL, exists := getDatabaseURL()
	if !exists {
		t.Fatal("must provide a valid TEST_DATABASE_URL environment variable")
	}

	conn, err := url.Parse(databaseURL)
	if err != nil {
		t.Fatalf("invalid database URL: %v", err)
	}

	pwd, _ := conn.User.Password()
	opts := &pg.Options{
		Addr:     conn.Host,
		Database: strings.Replace(conn.RequestURI(), "/", "", 1),
		User:     conn.User.Username(),
		Password: pwd,
	}

	pg, err := Connect(opts, app.StartDate)
	if err != nil {
		return nil, err
	}

	return pg, nil
}
