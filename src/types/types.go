package types

import "github.com/bwmarrin/discordgo"

type Command struct {
	Name                     string
	Description              string
	Category                 string
	Handler                  func(s *discordgo.Session, i *discordgo.InteractionCreate)
	DefaultMemberPermissions *int64
}

var (
	RegisteredCommands = make(map[string]*Command)
)
