package main

import (
	"fmt"
	"log"
	"os"
	"presenze-ectm-telegram-bot/pkg/db"
	"presenze-ectm-telegram-bot/pkg/repository"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

type score struct {
	wins, draws, losses uint
}

type application struct {
	client *tbot.Client
	score
	DSN string
	DB  repository.DatabaseRepository
}

var (
	app   application
	bot   *tbot.Server
	token string
	// options = map[string]string{
	// 	// choice : beats
	// 	"paper":    "rock",
	// 	"rock":     "scissors",
	// 	"scissors": "paper",
	// }
)

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	token = os.Getenv("TELEGRAM_TOKEN")
	log.Println("Token: ", token)

	app.DSN = os.Getenv("DSN")
}

func main() {
	connRDBMS, err := app.connectToDB()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error %s", err))
	}
	defer connRDBMS.Close()

	app.DB = &db.PostgresDBRepo{DB: connRDBMS}

	bot = tbot.New(token, tbot.WithWebhook("https://presenze-ectm.dellechiaie.it", ":4444"))
	// bot = tbot.New(token)
	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	bot.HandleMessage("/insert", app.insertHandler)
	// bot.HandleMessage("/score", app.scoreHandler)
	// bot.HandleMessage("/reset", app.resetHandler)
	bot.HandleCallback(app.callbackHandler)
	log.Fatal(bot.Start())
}
