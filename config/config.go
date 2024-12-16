package config

import (
	"fmt"
	"github.com/pkg/errors"
	pkgConfig "github.com/rishu/microservice/pkg/config"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

var (
	once   sync.Once
	config *Config
	err    error

	_, b, _, _ = runtime.Caller(0)
)

func Load() (*Config, error) {
	once.Do(func() {
		config, err = loadConfig()
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}
	return config, err
}

func TestConfigDir() string {
	configPath := filepath.Join(b, "..")
	return configPath
}

func loadConfig() (*Config, error) {

	conf := &Config{}
	testConfigDir := TestConfigDir()
	// loads config from file
	k, _, err := pkgConfig.LoadConfig(testConfigDir, "test")
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	err = k.UnmarshalWithConf("", conf, pkgConfig.DefaultUnmarshallingConfig(conf))
	if err != nil {
		return nil, fmt.Errorf("failed to refresh config: %w", err)
	}
	return conf, nil
}

type Config struct {
	MongoConfig *MongoConfig
	Server      *Server
}

type Server struct {
	Port         int
	GrpcPort     int
	GrpcHttpPort int
}

type MongoConfig struct {
	MongoConnectTimeoutMS time.Duration
	MongoDBName           string
	MongoDBURI            string
	MongoHost             string
	MongoMaxIdleTimeMS    time.Duration
	MongoMaxPoolSize      int
	MongoMinPoolSize      int
	MongoPassword         string
	MongoPort             int
}
