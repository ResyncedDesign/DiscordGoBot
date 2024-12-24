package types

import "github.com/bwmarrin/discordgo"

type Command struct {
	Name                     string
	Description              string
	Category                 string
	Handler                  func(s *discordgo.Session, i *discordgo.InteractionCreate)
	DefaultMemberPermissions *int64
	Options                  []*discordgo.ApplicationCommandOption
}

type Event struct {
    Name        string
    Description string
    Category    string
    Type        string
    Handler     interface{}
}

var (
	RegisteredCommands = make(map[string]*Command)
	RegisteredEvents   = make(map[string]*Event)
)
