package utils

import (
	"github.com/bwmarrin/discordgo"
)

func GetGuildCount(s *discordgo.Session) (int, error) {
	var allGuilds []*discordgo.UserGuild
	lastID := ""
	for {
		guilds, err := s.UserGuilds(200, "", lastID, false)
		if err != nil {
			return 0, err
		}

		allGuilds = append(allGuilds, guilds...)

		if len(guilds) < 200 {
			break
		}

		lastID = guilds[len(guilds)-1].ID
	}

	return len(allGuilds), nil
}
