package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	botApiKey := os.Getenv("TELEGRAM_EXAMPLE_BOT_API_KEY")
	if botApiKey == "" {
		fmt.Println("Error: TELEGRAM_EXAMPLE_BOT_API_KEY environment variable is not set or empty.")
		os.Exit(1)
	}

	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}


	b, err := bot.New(botApiKey, opts...)
	if err != nil {
		panic(err)
	}

	log.Println("Bot connected successfully")

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}