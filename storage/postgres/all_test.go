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
