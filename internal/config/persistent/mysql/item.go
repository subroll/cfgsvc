package mysql

import (
	"context"
	"database/sql"

	"github.com/subroll/cfgsvc/internal/config/persistent"
)

type itemStore struct {
	db *sql.DB
}

func (is *itemStore) Items(ctx context.Context, groupID, itemID int) ([]persistent.Item, error) {
	return nil, nil
}

func (is *itemStore) Insert(ctx context.Context, item persistent.Item) (persistent.Item, error) {
	return item, nil
}

func (is *itemStore) Update(ctx context.Context, item persistent.Item) (persistent.Item, error) {
	return item, nil
}

func (is *itemStore) Delete(ctx context.Context, id int) error {
	return nil
}
