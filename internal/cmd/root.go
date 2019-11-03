package cmd

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/subroll/cfgsvc/internal/app"
)

var rootCMD = &cobra.Command{
	Use:     app.Name,
	Short:   app.ShortDesc,
	Version: app.Version,
}

func init() {
	rootCMD.PersistentFlags().StringP("config", "", "file:config.yaml",
		`configuration requires <type:value>, currently the supported types are 'file', and so the value should be file path`)

	err := viper.BindPFlag("config_location", rootCMD.PersistentFlags().Lookup("config"))
	if err != nil {
		log.WithError(err).Fatal("could not bind config flag to config_location")
	}

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	conf := strings.Split(viper.GetString("config_location"), ":")
	if len(conf) != 2 {
		log.Fatal("configuration requires type:value")
	}

	err := app.LoadConfiguration(conf[0], conf[1])
	if err != nil {
		log.WithError(err).Fatal("fail to load configuration")
	}
}
