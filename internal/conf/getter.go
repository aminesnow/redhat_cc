package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	logLevel = "LOG_LEVEL"

	// service params
	serviceHost = "SERVICE_HOST"
	servicePort = "SERVICE_PORT"

	// storage type.
	// possible values: psql, memory
	storageType         = "STORAGE_TYPE"
	STORAGE_TYPE_PSQL   = "psql"
	STORAGE_TYPE_MEMORY = "memory"

	// postgresql params
	pgsqlHost   = "POSTGRESQL_HOST"
	pgsqlPort   = "POSTGRESQL_PORT"
	pgsqlUser   = "POSTGRESQL_USER"
	pgsqlPwd    = "POSTGRESQL_PASSWORD"
	pgsqlDBName = "POSTGRESQL_DBNAME"
)

// ParseConfiguration reads the configuration file given as parameter
func ParseConfiguration(confFile string) {

	setDefaults()

	viper.AutomaticEnv()

	if len(confFile) == 0 {
		logrus.Warn("No configuration file is defined")
	}

	viper.SetConfigFile(confFile)

	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Errorf("failed to read config file %v", confFile)
	}

	logrus.Infof("using config file: %v", viper.ConfigFileUsed())
}

func setDefaults() {
	viper.SetDefault(serviceHost, "localhost")
	viper.SetDefault(servicePort, 8080)

	viper.SetDefault(storageType, STORAGE_TYPE_MEMORY)

	viper.SetDefault(pgsqlHost, "localhost")
	viper.SetDefault(pgsqlPort, 5432)
	viper.SetDefault(pgsqlUser, "object_store_admin")
	viper.SetDefault(pgsqlDBName, "object_store")
}

func GetLogLevel() logrus.Level {
	l := viper.GetString(logLevel)

	level, err := logrus.ParseLevel(l)
	if err != nil {
		logrus.Warnf("unknown log level: %s", l)
		return logrus.DebugLevel
	}

	return level
}

func GetStorageType() string {
	st := viper.GetString(storageType)

	switch st {
	case STORAGE_TYPE_PSQL:
		return st
	case STORAGE_TYPE_MEMORY:
		return st
	default:
		return STORAGE_TYPE_MEMORY
	}
}

type ServiceParams struct {
	Host string
	Port int
}

func GetServiceParams() ServiceParams {
	return ServiceParams{
		Host: viper.GetString(serviceHost),
		Port: viper.GetInt(servicePort),
	}
}

type PsqlParams struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DBName string
}

func GetPsqlParams() PsqlParams {
	return PsqlParams{
		Host:   viper.GetString(pgsqlHost),
		Port:   viper.GetInt(pgsqlPort),
		User:   viper.GetString(pgsqlUser),
		Pwd:    viper.GetString(pgsqlPwd),
		DBName: viper.GetString(pgsqlDBName),
	}
}
