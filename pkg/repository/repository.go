package repository

import (
	"database/sql"
	"presenze-ectm-telegram-bot/pkg/data"
)

type DatabaseRepository interface {
	Connection() *sql.DB
	GetUserByTelegramUsername(telegramUsername string) (*data.User, error)
	GetUserByCognomeNome(firstName, lastName string) (*data.User, error)
	InsertPresence(presence data.Presence) error
}
