package server

import (
	"context"
	"time"

	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/danielnegri/jurassic-park-go/pkg/app"
	"github.com/danielnegri/jurassic-park-go/pkg/errors"
	"github.com/danielnegri/jurassic-park-go/pkg/guid"
	"github.com/danielnegri/jurassic-park-go/pkg/log"
	"github.com/danielnegri/jurassic-park-go/pkg/net"
	"github.com/danielnegri/jurassic-park-go/pkg/version"
	"github.com/danielnegri/jurassic-park-go/storage"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Run() error
	Shutdown()
}

type Config struct {
	HTTPServerConfig net.HTTPServerConfig

	// Set gin mode to release.
	ReleaseMode bool

	Storage storage.Storage

	// If specified, the server will use this function for determining time.
	Now func() time.Time
}

type service struct {
	cfg     Config
	guid    *guid.Generator
	health  gosundheit.Health
	logger  logrus.FieldLogger
	server  net.Server
	storage storage.Storage

	now func() time.Time
}

var _ Server = (*service)(nil)

func New(cfg Config) *service {
	if cfg.Now == nil {
		cfg.Now = time.Now
	}

	healthChecker := gosundheit.New()

	guid := guid.New(guid.Settings{
		StartTime: app.StartDate(),
	})

	svc := &service{
		cfg:     cfg,
		guid:    guid,
		health:  healthChecker,
		logger:  log.WithField("component", "server"),
		storage: cfg.Storage,
		now:     cfg.Now,
	}

	svc.server = net.NewServer(cfg.HTTPServerConfig, svc.newHandler(), svc.Shutdown)

	return svc
}

func (s *service) Run() error {
	const op errors.Op = "server.Run"
	s.logger.Infof("%s: Starting HTTP Server (%s)", app.Description, version.Version)

	ctx := context.Background()

	if s.storage != nil {
		if err := s.storage.Check(ctx); err != nil {
			s.logger.Errorf("error while checking connection with storage: %v", err)
		}
	} else {
		return errors.E(op, errors.KindUnexpected, "invalid storage configuration")
	}

	// Start Server
	if err := s.server.Run(); err != nil {
		return errors.E(op, "failed to start server", err)
	}

	return nil
}

func (s *service) Shutdown() {
	s.logger.Infof("%s: Stopping HTTP Server", app.Description)
	s.storage.Close()

}
