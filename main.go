package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"github.com/shomali11/slacker"
)

// getting channels to to prevent deadlock and mutex
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	// for each loop 
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3957073649254-3949140182919-E9jTNzdH9a4a8F52p3tQLoAy")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03TX3G46A3-3957083751990-1e2027dfab262ff74254417ce6115af5e4fa9e6a4cb1c14df96937d3f3c48f91")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	// <year> is the parameter
	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "Year of birth calculator",
		Examples: [] string {"My yob is 25", "My yob is 21"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := time.Now().Year() - yob

			//have to make a string r to input age {int} as Go is statically typed language
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}