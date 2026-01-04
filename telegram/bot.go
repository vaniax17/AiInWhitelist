package telegram

import (
	"aiInWhitelists/gemini"
	"context"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var waitingForAsk = make(map[int64]bool)

func InitTelegramBot() {
	token := os.Getenv("TELEGRAM_TOKEN")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(askHandler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		log.Fatal(err)
	}

	b.Start(ctx)
}

func askHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	userId := update.Message.From.ID
	text := strings.TrimSpace(strings.ToLower(update.Message.Text))

	if text == "ask" {
		waitingForAsk[userId] = true

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "type ur question",
		})
		if err != nil {
			return
		}
		return
	}

	if waitingForAsk[userId] {
		delete(waitingForAsk, userId)

		answer, err := gemini.Ask(update.Message.Text)
		if err != nil {
			log.Println(err)
			_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "ask failed",
			})
			return
		}

		_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   answer,
		})
	}
}
