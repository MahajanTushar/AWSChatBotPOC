package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shomali11/slacker"
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

func main() {
	token := "xoxb-1427234271280-1409918848549-bpSttkJiwUSFchKQHjjOPdj8"

	bot := slacker.NewClient(token)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("Hey I am <name>", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			name := "Hi" + request.Param("name") + "How may I help you today"
			response.Reply(name)
		},
	})

	bot.Command("tell card numbers associated to my account", &slacker.CommandDefinition{
		Description: "Echo a word!",
		Example:     "echo hello",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			word := "sure, please help me with your customer Id"
			response.Reply(word)
		},
	})

	bot.Command("acc no : <account-number>", &slacker.CommandDefinition{
		Description: "Echo a word!",
		Example:     "echo hello",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			cards := "410086CurrCon90, 510045Curcon23"
			response.Reply(cards)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
