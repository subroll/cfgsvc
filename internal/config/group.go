package config

import (
	"context"

	"github.com/subroll/cfgsvc/internal/config/persistent"
)

// GroupServer is the struct for group service layer
type GroupServer interface {
	Groups(ctx context.Context, id []int) ([]Group, error)
	Create(ctx context.Context, group Group) (Group, error)
	Change(ctx context.Context, group Group) (Group, error)
	Remove(ctx context.Context, itemID int) error
}

// NewGroup creates new group service layer
func NewGroup(gs persistent.GroupStore) GroupServer {
	return &groupService{groupStore: gs}
}

type groupService struct {
	groupStore persistent.GroupStore
}

func (gs *groupService) Groups(ctx context.Context, ids []int) ([]Group, error) {
	return nil, nil
}

func (gs *groupService) Create(ctx context.Context, group Group) (Group, error) {
	return group, nil
}

func (gs *groupService) Change(ctx context.Context, group Group) (Group, error) {
	return group, nil
}

func (gs *groupService) Remove(ctx context.Context, id int) error {
	return nil
}
