package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/subroll/cfgsvc/internal/config/persistent"
)

type groupStore struct {
	db *sql.DB
}

func (gs *groupStore) Groups(ctx context.Context, ids []int) ([]persistent.Group, error) {
	var args []interface{}
	condition := ""

	if len(ids) > 0 {
		idsParam := strings.TrimSuffix(strings.Repeat("?,", len(ids)), ",")
		condition = fmt.Sprintf(" WHERE id IN(%s)", idsParam)

		for _, id := range ids {
			args = append(args, id)
		}
	}

	baseQuery := `
		SELECT  
			id, 
			group_name, 
			created_at, 
			updated_at
		FROM 
			mst_config_group
		%s`
	q := fmt.Sprintf(baseQuery, condition)

	rows, err := gs.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		resultGroupID        sql.NullInt64
		resultGroupName      sql.NullString
		resultGroupCreatedAt sql.NullTime
		resultGroupUpdatedAt sql.NullTime
		groups               []persistent.Group
	)
	for rows.Next() {
		if err := rows.Scan(
			&resultGroupID,
			&resultGroupName,
			&resultGroupCreatedAt,
			&resultGroupUpdatedAt); err != nil {
			return nil, err
		}

		groups = append(groups, persistent.Group{
			ID:        int(resultGroupID.Int64),
			Name:      resultGroupName.String,
			CreatedAt: resultGroupCreatedAt.Time,
			UpdatedAt: resultGroupUpdatedAt.Time,
		})
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (gs *groupStore) Insert(ctx context.Context, group persistent.Group) (persistent.Group, error) {
	q := "INSERT INTO mst_config_group (group_name) VALUES (?)"

	result, err := gs.db.ExecContext(ctx, q, group.Name)
	if err != nil {
		return group, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return group, err
	}

	group.ID = int(id)

	return group, nil
}

func (gs *groupStore) Update(ctx context.Context, group persistent.Group) (persistent.Group, error) {
	q := "UPDATE mst_config_group SET group_name=? WHERE id=?"

	result, err := gs.db.ExecContext(ctx, q, group.Name, group.ID)
	if err != nil {
		return group, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return group, err
	}

	return group, nil
}

func (gs *groupStore) Delete(ctx context.Context, id int) error {
	q := `DELETE FROM mst_config_group WHERE id=?`

	result, err := gs.db.ExecContext(ctx, q, id)
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
