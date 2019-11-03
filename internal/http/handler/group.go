package handler

import (
	"net/http"

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
}

// Post is the http handler for /group
func (g *Group) Post(rw http.ResponseWriter, r *http.Request) {
}

// Put is the http handler for /group
func (g *Group) Put(rw http.ResponseWriter, r *http.Request) {
}

// Delete is the http handler for /group
func (g *Group) Delete(rw http.ResponseWriter, r *http.Request) {
}
