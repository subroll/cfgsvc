package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCMD = &cobra.Command{
	Use:   "serve",
	Short: "Start configuration http server",
	Run:   runServer,
}

func init() {
	serveCMD.PersistentFlags().IntP("port", "p", 8080, "port for the HTTP server")

	err := viper.BindPFlag("http.server.port", serveCMD.PersistentFlags().Lookup("port"))
	if err != nil {
		log.WithError(err).Fatal("could not bind port flag to http.server.port")
	}

	rootCMD.AddCommand(serveCMD)
}

func runServer(_ *cobra.Command, _ []string) {
}
