package messages

import (
	"DiscordGoBot/src/types"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func init() {
	types.RegisteredEvents["messageCreate"] = &types.Event{
		Name:        "messageCreate",
		Description: "Handles new messages",
		Category:    "messages",      // messages, guild, etc...
		Type:        "messageCreate", // messageCreate, GuildMemberAdd, etc...
		Handler:     handleMessageCreate,
	}
}

func handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content) // Contains the yap
}
