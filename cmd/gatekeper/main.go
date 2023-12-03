package main

import (
	"bot/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)
	producService := product.NewServise()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, producService)
		default:
			defaultBegavior(bot, update.Message)
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputmessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputmessage.Chat.ID, "/help - help\n"+
		"/list - list products",
	)

	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, producServise *product.Service) {
	outputMsgText := "Here all the products \n\n"

	product := producServise.List()
	for _, p := range product {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	bot.Send(msg)
}

func defaultBegavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	bot.Send(msg)
}
