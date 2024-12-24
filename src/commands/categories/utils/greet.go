package utils

import (
	"DiscordGoBot/src/types"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	permission := int64(discordgo.PermissionSendMessages)
	types.RegisteredCommands["greet"] = &types.Command{
		Name:                     "greet",
		Description:              "Greet a user",
		Category:                 "utils",
		Handler:                  handleGreet,
		DefaultMemberPermissions: &permission,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "user",
				Description: "The name of the user you want to greet",
				Required:    true,
			},
		},
	}
}

func handleGreet(s *discordgo.Session, i *discordgo.InteractionCreate) {
	username := i.ApplicationCommandData().Options[0].StringValue()

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Hello! %s", username),
		},
	})
	if err != nil {
		log.Printf("Error responding to greet command: %v", err)
	}
}
