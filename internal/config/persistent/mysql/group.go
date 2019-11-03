package mysql

import (
	"context"
	"database/sql"

	"github.com/subroll/cfgsvc/internal/config/persistent"
)

type groupStore struct {
	db *sql.DB
}

func (gs *groupStore) Groups(ctx context.Context, ids []int) ([]persistent.Group, error) {
	return nil, nil
}

func (gs *groupStore) Insert(ctx context.Context, group persistent.Group) (persistent.Group, error) {
	return persistent.Group{}, nil
}

func (gs *groupStore) Update(ctx context.Context, group persistent.Group) (persistent.Group, error) {
	return persistent.Group{}, nil
}

func (gs *groupStore) Delete(ctx context.Context, id int) error {
	return nil
}
