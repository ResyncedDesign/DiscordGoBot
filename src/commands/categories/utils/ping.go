package utils

import (
	"DiscordGoBot/src/types"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func init() {
	permission := int64(discordgo.PermissionSendMessages)
	types.RegisteredCommands["ping"] = &types.Command{
		Name:                     "ping",
		Description:              "Displays the bot's latency.",
		Category:                 "utils",
		Handler:                  handlePing,
		DefaultMemberPermissions: &permission,
	}
}

func handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	ID, err := strconv.ParseUint(i.ID, 10, 64)
	if err != nil {
		log.Printf("Error parsing interaction ID: %v", err)
		return
	}

	now := time.Now()
	interactionTimestamp := time.UnixMilli(int64(ID>>22) + 1420070400000)
	latency := now.Sub(interactionTimestamp)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Pong! Latency: %dms", latency.Milliseconds()),
		},
	})
	if err != nil {
		log.Printf("Error responding to ping command: %v", err)
	}
}
