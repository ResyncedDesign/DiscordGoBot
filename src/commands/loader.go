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

	log.Printf("Deleting %d guild commands\n", len(commands))

	for _, cmd := range commands {
		err := s.ApplicationCommandDelete(s.State.User.ID, guildID, cmd.ID)
		if err != nil {
			log.Printf("Failed to delete command %s: %v", cmd.ID, err)
		}
	}

	globalCommands, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		return err
	}

	log.Printf("Deleting %d global commands\n", len(globalCommands))

	for _, cmd := range globalCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
		if err != nil {
			log.Printf("Failed to delete global command %s: %v", cmd.ID, err)
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
			Options:                  cmd.Options,
		}

		if cmd.Global {
			_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command)
			if err != nil {
				return fmt.Errorf("error registering command %s: %w", cmd.Name, err)
			}
			log.Printf("Registered slash command: /%s (Category: %s, Global: %v)", cmd.Name, cmd.Category, cmd.Global)
			continue
		}

		_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, command)
		if err != nil {
			return fmt.Errorf("error registering command %s: %w", cmd.Name, err)
		}
		log.Printf("Registered slash command: /%s (Category: %s, Global: %v)", cmd.Name, cmd.Category, cmd.Global)
	}
	return nil
}
