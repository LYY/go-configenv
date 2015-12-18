package configenv

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type ConfigEnv struct {
	configFile  *yaml.File
	environment string
}

func NewEnv(configFile, environment string) *ConfigEnv {
	if environment != "" {
		names := strings.Split(configFile, ".")
		i := len(names) - 1
		names = append(names[:i], append([]string{environment}, names[i:]...)...)
		configFile = strings.Join(names, ".")
	}
	env := &ConfigEnv{
		configFile:  yaml.ConfigFile(configFile),
		environment: environment,
	}

	if env.configFile == nil {
		panic("go-configenv failed to open configFile: " + configFile)
	}

	return env
}

func (env *ConfigEnv) Get(spec, defaultValue string) string {
	value, err := env.configFile.Get(spec)
	if err != nil {
		value = defaultValue
	}
	return value
}

func (env *ConfigEnv) GetInt(spec string, defaultValue int) int {
	str := env.Get(spec, "")
	if str == "" {
		return defaultValue
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		log.Panic("go-configenv GetInt failed Atoi", env.environment, spec, str)
	}
	return val
}

func (env *ConfigEnv) GetDuration(spec string, defaultValue string) time.Duration {
	str := env.Get(spec, "")
	if str == "" {
		str = defaultValue
	}
	duration, err := time.ParseDuration(str)
	if err != nil {
		log.Panic("go-configenv GetDuration failed ParseDuration", env.environment, spec, str)
	}
	return duration
}

func (env *ConfigEnv) Require(spec string) string {
	value := env.Get(spec, "")
	if value == "" {
		log.Panicf("go-configenv Require couldn't find %s.%s", env.environment, spec)
	}
	return value
}

func (env *ConfigEnv) RequireInt(spec string) int {
	str := env.Require(spec)
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Panic("go-configenv RequireInt failed Atoi", env.environment, spec, str)
	}
	return val
}

func (env *ConfigEnv) RequireDuration(spec string) time.Duration {
	str := env.Require(spec)
	duration, err := time.ParseDuration(str)
	if err != nil {
		log.Panic("go-configenv RequireDuration failed ParseDuration", env.environment, spec, str)
	}
	return duration
}

func (env *ConfigEnv) Count(spec string) int {
	count, err := env.configFile.Count(spec)
	if err != nil {
		log.Panicf("go-configenv Count failed %s", err)
	}
	return count
}

func (env *ConfigEnv) GetList(spec string, defaultValue []string) []string {
	value, err := yamlGetList(env.configFile, spec)
	if err != nil {
		value = defaultValue
	}
	return value
}

func (env *ConfigEnv) GetEnvName() string {
	return env.environment
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	return value
}
