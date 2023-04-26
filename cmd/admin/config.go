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
	"net/url"
	"strings"

	"github.com/danielnegri/jurassic-park-go/pkg/log"
	"github.com/danielnegri/jurassic-park-go/pkg/net"
	"github.com/danielnegri/jurassic-park-go/server"
	"github.com/danielnegri/jurassic-park-go/storage"
	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func newLogger() logrus.FieldLogger {
	return log.New(viper.GetString("log_level"), viper.GetString("log_format"))
}

func newServerConfig(storage storage.Storage) server.Config {
	return server.Config{
		HTTPServerConfig: net.HTTPServerConfig{
			Addr: viper.GetString("addr"),
		},
		ReleaseMode: viper.GetString("log_level") != "debug",
		Storage:     storage,
	}
}

func newPostgresOptions() (*pg.Options, error) {
	databaseURL := viper.GetString("database")
	if databaseURL == "" {
		return nil, fmt.Errorf("invalid database URL")
	}

	conn, err := url.Parse(databaseURL)
	if err != nil {
		return nil, err
	}

	pwd, _ := conn.User.Password()
	db := &pg.Options{
		Addr:       conn.Host,
		Database:   strings.Replace(conn.RequestURI(), "/", "", 1),
		User:       conn.User.Username(),
		Password:   pwd,
		MaxConnAge: viper.GetDuration("database_max_conn_age"),
		PoolSize:   viper.GetInt("database_pool_size"),
	}

	return db, nil
}
