// Package handler provides the http handler.
package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/subroll/cfgsvc/internal/config"

	log "github.com/sirupsen/logrus"
)

var (
	badRequestResponse = response{
		HTTPCode: http.StatusBadRequest,
		Message:  "Bad Request",
	}
	notFoundResponse = response{
		HTTPCode: http.StatusNotFound,
		Message:  "Not Found",
	}
	internalServerErrorResponse = response{
		HTTPCode: http.StatusInternalServerError,
		Message:  "Internal Server Error",
	}
)

type response struct {
	HTTPCode int         `json:"-"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type itemsResponse struct {
	GroupID   int       `json:"group_id"`
	GroupName string    `json:"group_name"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Items     []struct {
		ID        int       `json:"id"`
		Key       string    `json:"key"`
		Value     string    `json:"value"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"items"`
}

func writeResponse(rw http.ResponseWriter, r response) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(r.HTTPCode)

	if err := json.NewEncoder(rw).Encode(r); err != nil {
		log.WithError(err).Error("fail to encode response to response writer")
	}
}

func makeGroupedItemsResponse(items []config.GroupedItems) response {
	confSlice := make([]itemsResponse, len(items))
	resp := response{
		HTTPCode: http.StatusOK,
		Data:     confSlice,
	}
	for i := 0; i < len(items); i++ {
		confSlice[i] = itemsResponse{
			GroupID:   items[i].GroupID,
			GroupName: items[i].GroupName,
			UpdatedAt: items[i].UpdatedAt,
			CreatedAt: items[i].CreatedAt,
		}
		for _, item := range items[i].Items {
			confSlice[i].Items = append(confSlice[i].Items, struct {
				ID        int       `json:"id"`
				Key       string    `json:"key"`
				Value     string    `json:"value"`
				UpdatedAt time.Time `json:"updated_at"`
				CreatedAt time.Time `json:"created_at"`
			}{
				ID:        item.ID,
				Key:       item.Key,
				Value:     item.Value,
				UpdatedAt: item.UpdatedAt,
				CreatedAt: item.CreatedAt,
			})
		}
	}

	return resp
}
