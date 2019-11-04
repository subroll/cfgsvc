// Package router provides the http router.
package router

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// NewRoute will create new http router
func NewRoute(
	getConfigHandler,
	postConfigHandler,
	putConfigHandler,
	deleteConfigHandler,
	getGroupHandler,
	postGroupHandler,
	putGroupHandler,
	deleteGroupHandler http.HandlerFunc,
	logMiddleware func(http.Handler) http.Handler) *mux.Router {
	root := mux.NewRouter()
	root.StrictSlash(true)
	root.Use(logMiddleware)

	root.Methods(http.MethodGet).Path("/").HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		_, err := rw.Write([]byte("hello..."))
		if err != nil {
			log.WithError(err).Error("can not write response")
		}
	})

	// config route
	root.Methods(http.MethodGet).Path("/config").HandlerFunc(getConfigHandler)
	root.Methods(http.MethodPost).Path("/config").HandlerFunc(postConfigHandler)
	root.Methods(http.MethodPut).Path("/config").HandlerFunc(putConfigHandler)
	root.Methods(http.MethodDelete).Path("/config").HandlerFunc(deleteConfigHandler)

	// group route
	root.Methods(http.MethodGet).Path("/group").HandlerFunc(getGroupHandler)
	root.Methods(http.MethodPost).Path("/group").HandlerFunc(postGroupHandler)
	root.Methods(http.MethodPut).Path("/group").HandlerFunc(putGroupHandler)
	root.Methods(http.MethodDelete).Path("/group").HandlerFunc(deleteGroupHandler)

	return root
}
