package handler

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

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
	ctx := r.Context()
	rawItemID := r.URL.Query().Get("id")
	rawGroupID := r.URL.Query().Get("group_id")
	var err error

	var itemID int
	if len(rawItemID) > 0 {
		itemID, err = strconv.Atoi(rawItemID)
		if err != nil {
			log.WithError(err).Error("can not convert id to int")
			writeResponse(rw, badRequestResponse)
			return
		}
		if itemID == 0 {
			log.Warn("id is zero")
			writeResponse(rw, notFoundResponse)
			return
		}
	}

	var groupID int
	if len(rawGroupID) > 0 {
		groupID, err = strconv.Atoi(rawGroupID)
		if err != nil {
			log.WithError(err).Error("can not convert group_id to int")
			writeResponse(rw, badRequestResponse)
			return
		}
		if groupID == 0 {
			log.Warn("group_id is zero")
			writeResponse(rw, notFoundResponse)
			return
		}
	}

	items, err := i.itemSvc.GroupedItems(ctx, groupID, itemID)
	if err != nil {
		if err == config.ErrNotFound {
			log.WithError(err).Warn("no config item matched the query")
			writeResponse(rw, notFoundResponse)
			return
		}

		log.WithError(err).Error("fail to get item list")
		writeResponse(rw, internalServerErrorResponse)
		return
	}

	writeResponse(rw, makeGroupedItemsResponse(items))
	return
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
