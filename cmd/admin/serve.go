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

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/danielnegri/jurassic-park-go/pkg/log"
	"github.com/danielnegri/jurassic-park-go/pkg/net"
	"github.com/danielnegri/jurassic-park-go/server"
	"github.com/danielnegri/jurassic-park-go/storage/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func commandServe() *cobra.Command {
	var (
		databaseURL        string
		databaseMaxConnAge time.Duration
		databasePoolSize   int
		logFormat          string
		logLevel           string
		addr               string
	)

	cmd := cobra.Command{
		Use:     "serve",
		Short:   "Start Ledger HTTP server",
		Example: fmt.Sprintf("%s serve", shortDescription),
		Run: func(cmd *cobra.Command, args []string) {
			if err := serve(); err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVar(&databaseURL, "database-url", "", "database connection string (ex.: postgresql://username:password@localhost/jurassic-park)")
	_ = viper.BindPFlag("database_url", cmd.Flags().Lookup("database-url"))

	cmd.Flags().DurationVar(&databaseMaxConnAge, "database-max-conn-age", postgres.DefaultMaxConnAge, "database connection age at which client retires (closes) the connection")
	_ = viper.BindPFlag("database_max_conn_age", cmd.Flags().Lookup("database-max-conn-age"))

	cmd.Flags().IntVar(&databasePoolSize, "database-pool-size", postgres.DefaultPoolSize(), "database maximum number of socket connections")
	_ = viper.BindPFlag("database_pool_size", cmd.Flags().Lookup("database-pool-size"))

	cmd.Flags().StringVar(&logFormat, "log-format", log.DefaultFormat, "logger format")
	_ = viper.BindPFlag("log_format", cmd.Flags().Lookup("log-format"))

	cmd.Flags().StringVar(&logLevel, "log-level", log.DefaultLevel, "logger level")
	_ = viper.BindPFlag("log_level", cmd.Flags().Lookup("log-level"))

	cmd.Flags().StringVar(&addr, "addr", net.DefaultAddr, "HTTP bind address")
	_ = viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))

	return &cmd
}

func serve() error {
	log.SetLogger(newLogger())

	pgOpts, err := newPostgresOptions()
	if err != nil {
		return err
	}

	now := time.Now

	pg, err := postgres.Connect(pgOpts, now)
	if err != nil {
		return err
	}

	cfg := newServerConfig(pg)
	cfg.Now = now

	s := server.New(cfg)
	if err := s.Run(); err != nil {
		return err
	}

	log.Infof("Finished")
	return nil
}
