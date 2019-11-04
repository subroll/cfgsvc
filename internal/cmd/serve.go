package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/subroll/cfgsvc/internal/app"
	"github.com/subroll/cfgsvc/internal/config"
	"github.com/subroll/cfgsvc/internal/config/persistent/mysql"
	"github.com/subroll/cfgsvc/internal/http/handler"
	"github.com/subroll/cfgsvc/internal/http/middleware"
	"github.com/subroll/cfgsvc/internal/http/router"
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
	storage, err := mysql.NewStorage()
	if err != nil {
		log.WithError(err).Fatal("fail to create storage layer")
	}
	defer func() {
		if err := storage.Close(); err != nil {
			log.WithError(err).Fatal("fail to close storage layer")
		}
	}()

	itemSvc := config.NewItem(storage.ItemStore, storage.GroupStore)
	itemHandler := handler.NewItem(itemSvc)

	groupSvc := config.NewGroup(storage.GroupStore)
	groupHandler := handler.NewGroup(groupSvc)

	httpServer := &http.Server{
		Addr: ":" + viper.GetString("http.server.port"),
		Handler: router.NewRoute(
			itemHandler.Get,
			itemHandler.Post,
			itemHandler.Put,
			itemHandler.Delete,
			groupHandler.Get,
			groupHandler.Post,
			groupHandler.Put,
			groupHandler.Delete,
			middleware.RequestLog),
	}

	idleConnClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.WithError(err).Error("http server shutdown error")
		}

		idleConnClosed <- struct{}{}
		close(idleConnClosed)
	}()

	log.WithFields(log.Fields{
		"port":        viper.GetString("http.server.port"),
		"app_version": app.Version,
	}).Info("starting http server")
	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.WithError(err).Error("fail to start the auth service")
	}

	<-idleConnClosed
	log.Info("http server gracefully stopped")
}
