package main

import (
	"fmt"
	"presenze-ectm-telegram-bot/pkg/data"
	"time"

	"github.com/yanzay/tbot/v2"
)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	msg := "Questo è un BOT della società ECTM e serve per segnalare la propria presenza al lavoro o la malattia"
	a.client.SendMessage(m.Chat.ID, msg)
	a.insertHandler(m)
}

// Handle the /insert command here
func (a *application) insertHandler(m *tbot.Message) {
	buttons := makeButtons()
	choosedOption := tbot.OptInlineKeyboardMarkup(buttons)
	// a.client.DeleteMessage(m.Chat.ID, m.MessageID)
	a.client.SendMessage(m.Chat.ID, "Scegli un'opzione:", choosedOption)
}

// // Handle the /score command here
// func (a *application) scoreHandler(m *tbot.Message) {
// 	msg := fmt.Sprintf("Scores:\nWins: %v\nDraws: %v\nLosses: %v", a.wins, a.draws, a.losses)
// 	a.client.SendMessage(m.Chat.ID, msg)
// }

// // Handle the /reset command here
// func (a *application) resetHandler(m *tbot.Message) {
// 	a.wins, a.draws, a.losses = 0, 0, 0
// 	msg := "Scores have been reset to 0."
// 	a.client.SendMessage(m.Chat.ID, msg)
// }

// Handle buttton presses here
func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	userChoice := cq.Data

	var user *data.User

	if cq.From.Username != "" {
		user, _ = a.DB.GetUserByTelegramUsername(cq.From.Username)
	} else {
		user, _ = a.DB.GetUserByCognomeNome(cq.From.LastName, cq.From.FirstName)
	}
	userID := user.Userid

	if userID > 0 {
		var presence = data.Presence{DataPresenza: time.Now().Format("2006-01-02"),
			Userid:       userID,
			FlagPresenza: userChoice,
		}
		a.DB.InsertPresence(presence)
	}

	// msg := a.draw(humanMove)
	msg := fmt.Sprintf("Registrata la: %s per l'Utente: %s, %s", userChoice, user.Cognome, user.Nome)
	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, msg)
}
