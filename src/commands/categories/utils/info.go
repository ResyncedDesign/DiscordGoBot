package utils

import (
	"DiscordGoBot/src/config"
	"DiscordGoBot/src/types"
	"fmt"
	"log"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
)

func init() {
	permission := int64(discordgo.PermissionSendMessages)
	types.RegisteredCommands["info"] = &types.Command{
		Name:                     "info",
		Description:              "Displays information about the bot",
		Category:                 "utils",
		Handler:                  handleInfo,
		DefaultMemberPermissions: &permission,
		Global:                   true,
	}
}

func handleInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Failed to retrieve build information")
		return
	}

	var allGuilds []*discordgo.UserGuild
	lastID := ""
	for {
		guilds, err := s.UserGuilds(200, "", lastID, false)
		if err != nil {
			log.Printf("Failed to get guild count: %v", err)
			return
		}

		allGuilds = append(allGuilds, guilds...)

		if len(guilds) < 200 {
			break
		}

		lastID = guilds[len(guilds)-1].ID
	}

	totalGuilds := len(allGuilds)

	embed := types.NewEmbed().
		SetTitle("✨ Made by resynced.design").
		AddField("⚙️ Go Version", info.GoVersion, true).
		AddField("⏰ Uptime", config.FormattedUptime(), true).
		AddField("🏰 Guild Count", fmt.Sprintf("%d", totalGuilds), true).
		SetImage("https://r2.resynced.design/cdn/01JFT00BNQ2R8K4DSVNZKY0R4H.png").
		SetColor(0x4e5454).MessageEmbed

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
	if err != nil {
		log.Printf("Failed to respond to interaction: %v", err)
	}
}
