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
	"context"
	"runtime"
	"strings"
	"time"

	"github.com/danielnegri/jurassic-park-go/pkg/errors"
	"github.com/danielnegri/jurassic-park-go/pkg/log"
	"github.com/danielnegri/jurassic-park-go/storage"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

const DefaultMaxConnAge = 10 * time.Minute

func DefaultPoolSize() int {
	return runtime.NumCPU() * 2
}

type Postgres struct {
	db     *pg.DB
	logger logrus.FieldLogger

	now func() time.Time
}

var _ storage.Storage = (*Postgres)(nil)

// Connect parses a database URL into options that can be used to connect to PostgreSQL.
func Connect(opts *pg.Options, now func() time.Time) (*Postgres, error) {
	const op errors.Op = "postgres.Connect"
	if now == nil {
		now = time.Now
	}

	if opts == nil {
		return nil, errors.E(op, errors.KindUnexpected, "invalid database config")
	}

	if opts.MaxConnAge == 0 {
		opts.MaxConnAge = DefaultMaxConnAge
	}

	if opts.PoolSize == 0 {
		opts.PoolSize = DefaultPoolSize()
	}

	logger := log.WithField("component", "postgres")
	logger.Infof("Connecting to %s/%s", opts.Addr, opts.Database)

	db := pg.Connect(opts)
	db.AddQueryHook(DebugHook{})

	return &Postgres{db: db, logger: logger, now: now}, nil
}

func (p *Postgres) Close() error {
	const op errors.Op = "postgres.Close"
	p.logger.Info("Closing database")

	if err := p.db.Close(); err != nil {
		return errors.E(op, errors.KindUnexpected, err)
	}

	return nil
}

func (p *Postgres) Check(ctx context.Context) error {
	const op errors.Op = "postgres.Check"

	if err := p.db.Ping(ctx); err != nil {
		return errors.E(op, errors.KindUnexpected, err)
	}

	return nil
}

func (p *Postgres) ExecTx(ctx context.Context, fn func(tx *pg.Tx) error) error {
	return p.db.RunInTransaction(ctx, fn)
}

func kind(err error) int {
	if strings.Contains(err.Error(), "duplicate key value") {
		return errors.KindAlreadyExists
	}

	if err == pg.ErrNoRows {
		return errors.KindNotFound
	}

	return errors.KindUnexpected
}
