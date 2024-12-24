package events

import (
	_ "DiscordGoBot/src/events/categories/core"
	_ "DiscordGoBot/src/events/categories/messages"
	"DiscordGoBot/src/types"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func LoadEvents() error {
	if len(types.RegisteredEvents) == 0 {
		return fmt.Errorf("no events were registered")
	}

	log.Printf("Registered %d events\n", len(types.RegisteredEvents))

	return nil
}

func RegisterEventHandlers(s *discordgo.Session) {
	for _, event := range types.RegisteredEvents {
		s.AddHandler(event.Handler)
	}
}
