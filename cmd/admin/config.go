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
