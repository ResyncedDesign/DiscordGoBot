package utils

import (
	"DiscordGoBot/src/config"
	"DiscordGoBot/src/types"
	"DiscordGoBot/src/utils"
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
	loadingEmbed := types.NewEmbed().
		SetTitle("Loading...").
		SetColor(0x4e5454).MessageEmbed

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{loadingEmbed},
		},
	})

	if err != nil {
		fmt.Println(err)
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Failed to retrieve build information")
		return
	}

	totalGuilds, err := utils.GetGuildCount(s)
	if err != nil {
		log.Printf("Failed to get guild count: %v", err)
		return
	}

	totalMembers, err := utils.GetTotalUserCount(s) // This takes a while. Ideally I would cache this using Redis or something similar.
	if err != nil {
		log.Printf("Failed to get guild user count: %v", err)
		return
	}

	embed := types.NewEmbed().
		SetTitle("✨ Made by resynced.design").
		AddField("⚙️ Go Version", info.GoVersion, true).
		AddField("⏰ Uptime", config.FormattedUptime(), true).
		AddField("🏰 Guild Count", fmt.Sprintf("%d", totalGuilds), true).
		AddField("👥 Member Count", fmt.Sprintf("%d", totalMembers), true).
		SetImage("https://r2.resynced.design/cdn/01JFT00BNQ2R8K4DSVNZKY0R4H.png").
		SetColor(0x4e5454).MessageEmbed

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Embeds: &[]*discordgo.MessageEmbed{embed},
		Components: &[]discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Emoji: &discordgo.ComponentEmoji{
							Name: "🌐",
						},
						Label: "Website",
						Style: discordgo.LinkButton,
						URL:   "https://resynced.design",
					},
					discordgo.Button{
						Emoji: &discordgo.ComponentEmoji{
							Name: "🔗",
						},
						Label: "Github",
						Style: discordgo.LinkButton,
						URL:   "https://github.com/ResyncedDesign",
					},
					discordgo.Button{
						Emoji: &discordgo.ComponentEmoji{
							Name: "📜",
						},
						Label: "Documentation",
						Style: discordgo.LinkButton,
						URL:   "https://docs.resynced.design/discord-go-bot/introduction",
					},
				},
			},
		},
	})
	if err != nil {
		log.Printf("Failed to respond to interaction: %v", err)
	}
}
