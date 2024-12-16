package config

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"time"
)

const (
	DefaultConfigDirectory = "config"
	configFileNameFormat   = "%s"
	ConfigType             = "yml"
)

var (
	DefaultUnmarshallingConfig = func(o interface{}) koanf.UnmarshalConf {
		return koanf.UnmarshalConf{
			DecoderConfig: &mapstructure.DecoderConfig{
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeDurationHookFunc(),
					mapstructure.StringToSliceHookFunc(","),
					mapstructure.StringToTimeHookFunc(time.RFC3339),
				),
				WeaklyTypedInput: true,
				Result:           o,
			},
		}
	}
)

func LoadConfig(testConfigDir, fileName string) (*koanf.Koanf, *string, error) {
	configDir, _ := GetConfigDir(testConfigDir)
	k, err := PopulateConfig(configDir, fileName)
	if err != nil {
		return nil, nil, errors.Wrap(err, "PopulateConfig failed")
	}
	return k, &configDir, nil
}

func PopulateConfig(configDir, fileName string) (*koanf.Koanf, error) {
	configPath := GetConfigPath(fileName, configDir)

	return populateConfigFromFiles(configPath)
}

func GetConfigPath(configName string, configDir string) string {
	fileName := fmt.Sprintf(configFileNameFormat, configName)
	return fmt.Sprintf("%s/%s.%s", configDir, fileName, ConfigType)
}

func populateConfigFromFiles(configPath string) (*koanf.Koanf, error) {
	// delimiter `.` used to read the config
	var k = koanf.New(".")

	// load YML config from file
	if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error loading config from config path: %s", configPath))
	}

	if err := k.Load(env.Provider("", ".", func(s string) string { return s }), nil); err != nil {
		return nil, errors.Wrap(err, "error loading env variables")
	}

	return k, nil
}

func GetConfigDir(testConfigDir string) (string, error) {
	env, _ := os.LookupEnv("ENV")
	if env == "test" {
		return testConfigDir, nil
	}

	configDir, ok := os.LookupEnv("CONFIG_DIR")
	if !ok {
		currDir, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("CONFIG_DIR not found")
		}
		configDir = filepath.Join(currDir, DefaultConfigDirectory)
	}
	return configDir, nil
}
