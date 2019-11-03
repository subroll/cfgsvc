// Package cmd provides the functionality of structured command line application.
package cmd

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		log.WithError(err).Fatal("fail to execute root command")
	}
}
