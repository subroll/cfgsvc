package handler

import (
	"net/http"

	"github.com/subroll/cfgsvc/internal/config"
)

// NewItem create new config item pointer and returned as Item pointer
func NewItem(itemSvc config.ItemServer) *Item {
	return &Item{itemSvc: itemSvc}
}

// Item is the structure of config item http handler
type Item struct {
	itemSvc config.ItemServer
}

// Get is the http handler for /config
func (i *Item) Get(rw http.ResponseWriter, r *http.Request) {
}

// Post is the http handler for /config
func (i *Item) Post(rw http.ResponseWriter, r *http.Request) {
}

// Put is the http handler for /config
func (i *Item) Put(rw http.ResponseWriter, r *http.Request) {
}

// Delete is the http handler for /config
func (i *Item) Delete(rw http.ResponseWriter, r *http.Request) {
}
