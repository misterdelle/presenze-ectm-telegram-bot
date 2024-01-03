package db

import (
	"context"
	"database/sql"
	"log"
	"presenze-ectm-telegram-bot/pkg/data"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

// GetUserByTelegramUsername returns a user by his Telegram Username
func (m *PostgresDBRepo) GetUserByTelegramUsername(telegramUsername string) (*data.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var user data.User

	query := `select user_id, nome, cognome, telegram_username
	            from "Presenze".utente
			   where telegram_username = $1`

	rows, err := m.DB.QueryContext(ctx, query, telegramUsername)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&user.Userid,
			&user.Nome,
			&user.Cognome,
			&user.TelegramUsername,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	return &user, nil
}

// GetUserByCognomeNome returns a user by his Telegram Username
func (m *PostgresDBRepo) GetUserByCognomeNome(lastName, firstName string) (*data.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var user data.User

	query := `select user_id, nome, cognome, telegram_username
	            from "Presenze".utente
			   where cognome = $1
			     and nome = $2`

	rows, err := m.DB.QueryContext(ctx, query, lastName, firstName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&user.Userid,
			&user.Nome,
			&user.Cognome,
			&user.TelegramUsername,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	return &user, nil
}
