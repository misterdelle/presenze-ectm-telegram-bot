package main

import (
	"fmt"
	"log"
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
	a.client.SendMessage(m.Chat.ID, "Scegli un'opzione:", choosedOption)
}

// Handle buttton presses here
func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	userChoice := cq.Data

	var user *data.User

	if cq.From.Username != "" {
		log.Printf("Username: %s - LastName: %s - FirstName: %s - UserChoice: %s\n", cq.From.Username, cq.From.LastName, cq.From.FirstName, userChoice)
		user, _ = a.DB.GetUserByTelegramUsername(cq.From.Username)
	} else {
		log.Printf("LastName: %s - FirstName: %s - UserChoice: %s\n", cq.From.LastName, cq.From.FirstName, userChoice)
		user, _ = a.DB.GetUserByCognomeNome(cq.From.LastName, cq.From.FirstName)
	}

	var msg string

	if user != nil {
		userID := user.Userid

		if userID > 0 {
			var presence = data.Presence{DataPresenza: time.Now().Format("2006-01-02"),
				Userid:       userID,
				FlagPresenza: userChoice,
			}
			a.DB.InsertPresence(presence)

			msg = fmt.Sprintf("Registrato: %s per l'Utente: %s, %s", userChoice, user.Cognome, user.Nome)
		} else {
			log.Printf("Username: %s - LastName: %s - FirstName: %s NOT FOUND\n", cq.From.Username, cq.From.LastName, cq.From.FirstName)
			msg = fmt.Sprintf("Utente: %s - LastName: %s - FirstName: %s NON TROVATO", cq.From.Username, cq.From.LastName, cq.From.FirstName)
		}
	}

	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, msg)
}
