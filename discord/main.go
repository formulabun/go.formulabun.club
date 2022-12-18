package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"GoBun/services/discord/context"
	"GoBun/services/discord/env"
	"GoBun/services/discord/slashplayers"
	"GoBun/services/discord/status"

	"github.com/bwmarrin/discordgo"
)


func main() {
	env.ValidateEnvironment()

	log.Print("Discord service is starting.")
	session, err := discordgo.New(fmt.Sprintf("Bot %s", env.TOKEN))
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not create a client: %s", err))
	}

	session.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
		log.Print("Ready")
	})

	err = session.Open()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not start client: %s", err))
	}

	ctx := context.DiscordContext{session, make(chan struct{})}

	go status.Start(ctx)
	go slashplayers.Start(ctx)

	waitForShutdown(ctx)
}

func waitForShutdown(c context.DiscordContext) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	<-signals

	log.Print("Closing")
	close(c.Cancel)
	c.S.Close()
}
