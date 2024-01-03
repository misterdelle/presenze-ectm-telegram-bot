package db

import (
	"context"
	"errors"
	"presenze-ectm-telegram-bot/pkg/data"
)

// InsertPresence inserts a new Presence into the database
func (m *PostgresDBRepo) InsertPresence(presence data.Presence) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into "Presenze".presenza (data_presenza, user_id, flag_presenza)
		values ($1, $2, $3)`

	row := m.DB.QueryRowContext(ctx, stmt,
		presence.DataPresenza,
		presence.Userid,
		presence.FlagPresenza)

	if row != nil {
		return errors.New("error inserting presence")
	}

	return nil
}
