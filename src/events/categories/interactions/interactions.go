package messages

import (
	"DiscordGoBot/src/types"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func init() {
	types.RegisteredEvents["interactionCreate"] = &types.Event{
		Name:        "interactionCreate",
		Description: "Handle interaction create events",
		Category:    "interactions",
		Type:        "interactionCreate",
		Handler:     handleInteractions,
	}
}

func handleInteractions(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionModalSubmit:
		data := i.ModalSubmitData()
		modalID := strings.Split(data.CustomID, "_")[0]

		for _, cmd := range types.RegisteredCommands {
			if cmd.Modal != nil && cmd.Name == modalID {
				cmd.Modal.Handler(s, i)
				return
			}
		}
		log.Printf("No handler found for modal: %s", modalID)
	default:
		log.Printf("Unhandled interaction type: %s", i.Type)
	}

}
