package core

import (
	"DiscordGoBot/src/commands"
	"DiscordGoBot/src/types"
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	types.RegisteredEvents["ready"] = &types.Event{
		Name:        "ready",
		Description: "Handles bot ready event",
		Category:    "core",
		Type:        "ready",
		Handler:     handleReady,
	}
}

func handleReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as: %v", s.State.User.Username)
	log.Println("Bot is now running. Press CTRL+C to exit.")

	if err := commands.DeleteAllCommands(s); err != nil {
		log.Printf("Error deleting slash commands: %v", err)
	}

	if err := commands.RegisterSlashCommands(s); err != nil {
		log.Printf("Error registering slash commands: %v", err)
	}
}
