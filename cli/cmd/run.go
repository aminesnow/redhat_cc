package cmd

import (
	"fmt"
	"os"

	"github.com/aminesnow/redhat_cc/internal/conf"
	"github.com/aminesnow/redhat_cc/internal/repo"
	"github.com/aminesnow/redhat_cc/internal/repo/memory"
	"github.com/aminesnow/redhat_cc/internal/repo/postgres"
	"github.com/aminesnow/redhat_cc/internal/usecases"
	"github.com/aminesnow/redhat_cc/restapi/server"
	"github.com/aminesnow/redhat_cc/restapi/setup"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	gorm_pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// runCmd represents the run command.
var runCmd = &cobra.Command{
	Use:          "run",
	Short:        "Run the service",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.SetLevel(logrus.DebugLevel)

		conf.ParseConfiguration(cfgFile)

		// setup logger
		logrus.SetOutput(os.Stdout)
		lvl := conf.GetLogLevel()
		logrus.SetLevel(lvl)

		// setup repo
		var repo repo.ObjectStore

		st := conf.GetStorageType()
		switch st {
		case conf.STORAGE_TYPE_MEMORY:
			repo = memory.NewMemoryObjectRepo()
		case conf.STORAGE_TYPE_PSQL:
			psqlConf := conf.GetPsqlParams()
			dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
				psqlConf.Host, psqlConf.User, psqlConf.Pwd, psqlConf.DBName, psqlConf.Port)
			db, err := gorm.Open(gorm_pg.Open(dsn), &gorm.Config{})
			if err != nil {
				logrus.WithError(err).Error("failed to connect to postgresql")
				return err
			}

			repo = postgres.NewPostgresqlRepo(db)
		default:
			repo = memory.NewMemoryObjectRepo()
		}

		// setup server
		uc := usecases.NewObjectManager(repo)
		api := setup.GetSwaggerAPI()

		setup.SetHandlers(api, uc)

		//api.ServerShutdown = apiServerShutDown(*producer)

		serverParams := conf.GetServiceParams()
		server := server.NewServer(api)

		server.Port = serverParams.Port
		server.Host = serverParams.Host

		server.SetHandler(setup.SetAPIMiddleware(api))

		if err := server.Serve(); err != nil {
			logrus.WithError(err).Error("failed to serve")

			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

/*
func apiServerShutDown() func() {
	return func() {
		// TODO
	}
}*/
