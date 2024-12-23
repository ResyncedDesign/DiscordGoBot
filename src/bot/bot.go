package bot

import (
	"DiscordGoBot/src/commands"
	"DiscordGoBot/src/types"
	"log"
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

	return &Bot{
		session: session,
	}, nil
}

func (b *Bot) Start() error {
	b.session.Identify.Intents = discordgo.IntentsGuildPresences |
		discordgo.IntentsGuildMembers |
		discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages

	b.session.AddHandler(b.ready)

	commands.RegisterCommandHandlers(b.session)

	return b.session.Open()
}

func (b *Bot) Close() {
	b.session.Close()
}

func (b *Bot) ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as: %v", s.State.User.Username)
	log.Println("Bot is now running. Press CTRL+C to exit.")

	if err := commands.DeleteAllCommands(s); err != nil {
		log.Printf("Error deleting slash commands: %v", err)
	}

	if err := commands.RegisterSlashCommands(s); err != nil {
		log.Printf("Error registering slash commands: %v", err)
	}
}
