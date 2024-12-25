package types

import "github.com/bwmarrin/discordgo"

type Command struct {
	Name                     string
	Description              string
	Category                 string
	Handler                  func(s *discordgo.Session, i *discordgo.InteractionCreate)
	Modal                    *Modal
	DefaultMemberPermissions *int64
	Options                  []*discordgo.ApplicationCommandOption
	Global                   bool // Whether the command should be registered globally or only in your guild
}

type Event struct {
	Name        string
	Description string
	Category    string
	Type        string
	Handler     interface{}
}

type Modal struct {
	ID         string
	Title      string
	Components []discordgo.MessageComponent
	Handler    func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var (
	RegisteredCommands = make(map[string]*Command)
	RegisteredEvents   = make(map[string]*Event)
)
