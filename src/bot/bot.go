package bot

import (
	"DiscordGoBot/src/commands"
	"DiscordGoBot/src/events"
	"DiscordGoBot/src/types"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
}

func New() (*Bot, error) {
	token := os.Getenv("TOKEN")
	if token == "" {
		return nil, types.ErrTokenNotFound
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	if err := commands.LoadCommands(); err != nil {
		return nil, err
	}

	if err := events.LoadEvents(); err != nil {
		return nil, err
	}

	return &Bot{
		session: session,
	}, nil
}

func (b *Bot) Start() error {
	b.session.Identify.Intents = discordgo.IntentsGuildPresences |
		discordgo.IntentsGuildMembers |
		discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages

	events.RegisterEventHandlers(b.session)
	commands.RegisterCommandHandlers(b.session)

	return b.session.Open()
}

func (b *Bot) Close() {
	b.session.Close()
}
