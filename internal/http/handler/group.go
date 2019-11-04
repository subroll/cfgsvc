package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/subroll/cfgsvc/internal/config"
)

// NewGroup create new group pointer and returned as Group pointer
func NewGroup(groupSvc config.GroupServer) *Group {
	return &Group{groupSvc: groupSvc}
}

// Group is the structure of config group http handler
type Group struct {
	groupSvc config.GroupServer
}

// Get is the http handler for /group
func (g *Group) Get(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rawIDS := r.URL.Query()["id"]

	var ids []int
	if len(rawIDS) > 0 {
		for _, rawID := range rawIDS {
			id, err := strconv.Atoi(rawID)
			if err != nil {
				log.WithError(err).Error("can not convert id to int")
				writeResponse(rw, badRequestResponse)
				return
			}
			if id == 0 {
				log.Warn("id is zero")
				writeResponse(rw, notFoundResponse)
				return
			}
			ids = append(ids, id)
		}
	}

	groups, err := g.groupSvc.Groups(ctx, ids)
	if err != nil {
		if err == config.ErrNotFound {
			log.WithError(err).Warn("no config group matched the query")
			writeResponse(rw, notFoundResponse)
			return
		} else if err == config.ErrInvalidGroup {
			log.WithError(err).Warn("one of more group id is invalid")
			writeResponse(rw, badRequestResponse)
			return
		}

		log.WithError(err).Error("fail to get group list")
		writeResponse(rw, internalServerErrorResponse)
		return
	}

	writeResponse(rw, makeGroupResponse(groups))
	return
}

// Post is the http handler for /group
func (g *Group) Post(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	b, err := readRequestBody(r)
	if err != nil {
		log.WithError(err).Error("fail to read request body")
		writeResponse(rw, badRequestResponse)
	}

	var gr createGroupRequest
	if err = json.Unmarshal(b, &gr); err != nil {
		log.WithError(err).Error("fail to unmarshal request body")
		writeResponse(rw, badRequestResponse)
		return
	}

	group, err := g.groupSvc.Create(ctx, config.Group{Name: gr.Name})
	if err != nil {
		if err == config.ErrRequiredField {
			log.WithError(err).Error("no config group created")
			writeResponse(rw, badRequestResponse)
			return
		}

		log.WithError(err).Error("fail to create config group")
		writeResponse(rw, internalServerErrorResponse)
		return
	}

	writeResponse(rw, makeGroupResponse([]config.Group{group}))
	return
}

// Put is the http handler for /group
func (g *Group) Put(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	b, err := readRequestBody(r)
	if err != nil {
		log.WithError(err).Error("fail to read request body")
		writeResponse(rw, badRequestResponse)
	}

	var gr updateGroupRequest
	if err = json.Unmarshal(b, &gr); err != nil {
		log.WithError(err).Error("fail to unmarshal request body")
		writeResponse(rw, badRequestResponse)
		return
	}

	group, err := g.groupSvc.Change(ctx, config.Group{ID: gr.ID, Name: gr.Name})
	if err != nil {
		if err == config.ErrInvalidGroup || err == config.ErrRequiredField {
			log.WithError(err).Error("no config group updated")
			writeResponse(rw, badRequestResponse)
			return
		}

		log.WithError(err).Error("fail to update config group")
		writeResponse(rw, internalServerErrorResponse)
		return
	}

	writeResponse(rw, makeGroupResponse([]config.Group{group}))
	return
}

// Delete is the http handler for /group
func (g *Group) Delete(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rawID := r.URL.Query().Get("id")
	var err error

	var id int
	if len(rawID) > 0 {
		id, err = strconv.Atoi(rawID)
		if err != nil {
			log.WithError(err).Error("can not convert id to int")
			writeResponse(rw, badRequestResponse)
			return
		}
		if id == 0 {
			log.Warn("id is zero")
			writeResponse(rw, badRequestResponse)
			return
		}
	}

	if err := g.groupSvc.Remove(ctx, id); err != nil {
		if err == config.ErrNoRecordRemoved || err == config.ErrRequiredField {
			log.WithError(err).Error("no config item removed")
			writeResponse(rw, badRequestResponse)
			return
		}

		log.WithError(err).Error("fail to remove config item")
		writeResponse(rw, internalServerErrorResponse)
		return
	}

	writeResponse(rw, response{
		HTTPCode: http.StatusOK,
		Data: struct {
			ID int `json:"id"`
		}{ID: id},
	})
}
