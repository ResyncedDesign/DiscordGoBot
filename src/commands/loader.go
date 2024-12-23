package commands

import (
	_ "DiscordGoBot/src/commands/categories/utils"
	"DiscordGoBot/src/types"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func LoadCommands() error {
	if len(types.RegisteredCommands) == 0 {
		return fmt.Errorf("no commands were registered")
	}

	return nil
}

func DeleteAllCommands(s *discordgo.Session) error {
	guildID := os.Getenv("GUILDID")
	if guildID == "" {
		return log.Output(0, types.ErrGuildIDNotFound.Error())
	}

	commands, err := s.ApplicationCommands(s.State.User.ID, guildID)
	if err != nil {
		return err
	}

	log.Printf("Deleting %d commands\n", len(commands))

	for _, cmd := range commands {
		err := s.ApplicationCommandDelete(s.State.User.ID, guildID, cmd.ID)
		if err != nil {
			log.Printf("Failed to delete command %s: %v", cmd.ID, err)
		}
	}

	return nil
}

func RegisterCommandHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		commandName := i.ApplicationCommandData().Name
		if command, exists := types.RegisteredCommands[commandName]; exists {
			command.Handler(s, i)
		}
	})
}

func RegisterSlashCommands(s *discordgo.Session) error {
	guildID := os.Getenv("GUILDID")
	if guildID == "" {
		return log.Output(0, types.ErrGuildIDNotFound.Error())
	}

	for _, cmd := range types.RegisteredCommands {
		command := &discordgo.ApplicationCommand{
			Name:                     cmd.Name,
			Description:              cmd.Description,
			DefaultMemberPermissions: cmd.DefaultMemberPermissions,
		}

		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, command)
		if err != nil {
			return fmt.Errorf("error registering command %s: %w", cmd.Name, err)
		}
		log.Printf("Registered slash command: /%s (Category: %s)", cmd.Name, cmd.Category)
	}
	return nil
}
