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
	}
}

func handleInfo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("Failed to retrieve build information")
		return
	}

	embed := types.NewEmbed().
		SetTitle("Bot Information").
		SetDescription("✨ Made by resynced.design").
		AddField("⚙️ Go Version", info.GoVersion, true).
		AddField("⏰ Uptime", config.FormattedUptime(), true).
		SetImage("https://r2.resynced.design/cdn/01JFB8A2RZRMCHH6RBYP32R40B.png"). // TODO: Change to a more appropriate image
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
