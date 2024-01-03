package main

import (
	"github.com/yanzay/tbot/v2"
)

// var picks = []string{"rock", "paper", "scissors"} // choices from where the bot picks

func makeButtons() *tbot.InlineKeyboardMarkup {
	// Create buttons with visible Text and CallbackData as a value.
	btnPresente := tbot.InlineKeyboardButton{
		Text:         "👍 Presente",
		CallbackData: "Presente",
	}
	btnMalattia := tbot.InlineKeyboardButton{
		Text:         "🚑 Malattia",
		CallbackData: "Malattia",
	}
	buttons := []tbot.InlineKeyboardButton{btnPresente, btnMalattia}

	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			buttons,
		},
	}
}

// func (a *application) draw(humanMove string) (msg string) {
// 	var result string
// 	botMove := picks[rand.Intn(len(picks))] // Generate a random option for the bot

// 	// Determine the outcome
// 	switch humanMove {
// 	case botMove:
// 		result = "drew"
// 		a.draws++
// 	case options[botMove]:
// 		result = "lost"
// 		a.losses++
// 	default:
// 		result = "won"
// 		a.wins++
// 	}
// 	msg = fmt.Sprintf("You %s! You chose %s and I chose %s.", result, humanMove, botMove)
// 	return
// }
