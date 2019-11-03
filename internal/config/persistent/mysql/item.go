package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/subroll/cfgsvc/internal/config/persistent"
)

type itemStore struct {
	db *sql.DB
}

func (is *itemStore) Items(ctx context.Context, groupID, itemID int) ([]persistent.Item, error) {
	var args []interface{}
	condition := ""

	if groupID > 0 || itemID > 0 {
		condition += " WHERE"
	}

	if groupID > 0 {
		args = append(args, groupID)
		condition += " item_group.id = ?"
	}

	if groupID > 0 && itemID > 0 {
		condition += " AND"
	}

	if itemID > 0 {
		args = append(args, itemID)
		condition += " item.id = ?"
	}

	baseQuery := `
		SELECT 
			item.id, 
			item.key, 
			item.value, 
			item.created_at, 
			item.updated_at, 
			item_group.id, 
			item_group.group_name, 
			item_group.created_at, 
			item_group.updated_at
		FROM 
			mst_config as item
		LEFT JOIN mst_config_group item_group 
			ON item.mst_config_group_id = item_group.id
		%s`
	q := fmt.Sprintf(baseQuery, condition)

	rows, err := is.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		resultItemID         sql.NullInt64
		resultItemKey        sql.NullString
		resultItemValue      sql.NullString
		resultItemCreatedAt  sql.NullTime
		resultItemUpdatedAt  sql.NullTime
		resultGroupID        sql.NullInt64
		resultGroupName      sql.NullString
		resultGroupCreatedAt sql.NullTime
		resultGroupUpdatedAt sql.NullTime
		items                []persistent.Item
	)
	for rows.Next() {
		if err := rows.Scan(
			&resultItemID,
			&resultItemKey,
			&resultItemValue,
			&resultItemCreatedAt,
			&resultItemUpdatedAt,
			&resultGroupID,
			&resultGroupName,
			&resultGroupCreatedAt,
			&resultGroupUpdatedAt); err != nil {
			return nil, err
		}

		items = append(items, persistent.Item{
			ID:        int(resultItemID.Int64),
			Key:       resultItemKey.String,
			Value:     resultItemValue.String,
			CreatedAt: resultItemCreatedAt.Time,
			UpdatedAt: resultItemUpdatedAt.Time,
			Group: persistent.Group{
				ID:        int(resultGroupID.Int64),
				Name:      resultGroupName.String,
				CreatedAt: resultGroupCreatedAt.Time,
				UpdatedAt: resultGroupUpdatedAt.Time,
			},
		})
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (is *itemStore) Insert(ctx context.Context, item persistent.Item) (persistent.Item, error) {
	q := "INSERT INTO mst_config (mst_config_group_id, `key`, `value`) VALUES (?,?,?)"

	result, err := is.db.ExecContext(ctx, q, item.Group.ID, item.Key, item.Value)
	if err != nil {
		return item, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return item, err
	}

	item.ID = int(id)

	return item, nil
}

func (is *itemStore) Update(ctx context.Context, item persistent.Item) (persistent.Item, error) {
	q := "UPDATE mst_config SET mst_config_group_id=?, `key`=?, `value`=? WHERE id=?"

	result, err := is.db.ExecContext(ctx, q, item.Group.ID, item.Key, item.Value, item.ID)
	if err != nil {
		return item, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return item, err
	}

	return item, nil
}

func (is *itemStore) Delete(ctx context.Context, id int) error {
	q := `DELETE FROM mst_config WHERE id=?`

	result, err := is.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	ar, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if ar <= 0 {
		return persistent.ErrNoRowsAffected
	}

	return nil
}
