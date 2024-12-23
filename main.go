package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"DiscordGoBot/src/bot"
	"DiscordGoBot/src/config"
)

func main() {
	config.Load()

	discordBot, err := bot.New()
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	err = discordBot.Start()
	if err != nil {
		log.Fatalf("Error starting bot: %v", err)
	}
	defer discordBot.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down the bot...")
}
