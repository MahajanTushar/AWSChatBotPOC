package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"errors"

	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Message)
		fmt.Println()
	}
}

func OnlyErrors() error {
	return errors.New("something went wrong!")
}

func main() {
	//lambda.Start(OnlyErrors)

	token := os.Getenv("botToken")

	bot := slacker.NewClient(token)

	go printCommandEvents(bot.CommandEvents())

	bot.Init(func() {
		fmt.Println(" token is " + token)
		log.Println("Connected!")
	})

	bot.Err(func(err string) {
		log.Println(err)
	})

	bot.Command("Hey I am <name>", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			name := "Hi " + request.Param("name") + " How may I help you today."
			response.Reply(name)
		},
	})

	bot.Command("tell card numbers associated to my account", &slacker.CommandDefinition{
		Description: "ask for customer Id when user asks for the card numbers",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			word := "sure, please help me with your customer Id"
			response.Reply(word)
		},
	})

	bot.Command("acc no : <account-number>", &slacker.CommandDefinition{
		Description: "send harcoded account numbers back all the time",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {

			cards := "410086CurrCon90, 510045Curcon23"

			response.Reply(cards)
		},
	})

	bot.Command(" more info ? ", &slacker.CommandDefinition{
		Description: "send the source where more info on this can be found",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			url := "kindly log on to : https://www.tsys.com"
			attachments := []slack.Attachment{}

			attachments = append(attachments, slack.Attachment{
				Color:    "red",
				ImageURL: "https://www.tsys.com",
				Actions: []slack.AttachmentAction{
					slack.AttachmentAction{
						URL: "https://www.google.com",
					},
				},
			})

			response.Reply(url, slacker.WithAttachments(attachments))
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
