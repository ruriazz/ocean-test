package manager

import (
	"context"
	"os"

	"github.com/ruriazz/ocean-test/pkg/configs"
	"github.com/ruriazz/ocean-test/pkg/db"
	"github.com/ruriazz/ocean-test/pkg/server"
	log "github.com/sirupsen/logrus"
)

type Manager interface {
	Log() *log.Logger
	Config() configs.Config
	Server() server.Server
	Redis() db.RedisClient
	StartServer()
}

type manager struct {
	log    *log.Logger
	config configs.Config
	server server.Server
	redis  db.RedisClient
}

func New(ctx context.Context) Manager {
	logger := newLog()

	config, err := configs.LoadConfig()
	if err != nil {
		logger.WithFields(log.Fields{
			"context": "Manager::configs.LoadConfig",
		}).Fatal(err)
	}

	server, err := server.New()
	if err != nil {
		logger.WithFields(log.Fields{
			"context": "Manager::server.New",
		}).Fatal(err)
	}

	redis, err := db.NewRedisClient(ctx, *config)
	if err != nil {
		logger.WithFields(log.Fields{
			"context": "Manager::db.New",
		}).Fatal(err)
	}

	return manager{
		log:    logger,
		config: *config,
		server: server,
		redis:  redis,
	}
}

func newLog() *log.Logger {
	logger := log.New()

	logger.Out = os.Stdout
	logger.SetFormatter(&log.TextFormatter{
		ForceQuote:    true,
		FullTimestamp: true,
	})

	return logger
}

func (m manager) Log() *log.Logger {
	return m.log
}

func (m manager) Config() configs.Config {
	return m.config
}

func (m manager) StartServer() {
	if err := m.server.Start(":8000"); err != nil {
		m.Log().WithFields(log.Fields{
			"context": "Manager::StartServer",
		}).Fatal(err)
	}
}

func (m manager) Server() server.Server {
	return m.server
}

func (m manager) Redis() db.RedisClient {
	return m.redis
}
