package commands

import (
	_ "DiscordGoBot/src/commands/categories/modals"
	_ "DiscordGoBot/src/commands/categories/utils"
	"DiscordGoBot/src/types"
	"fmt"
	"log"
	"os"
	"strings"

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
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			commandName := i.ApplicationCommandData().Name
			if command, exists := types.RegisteredCommands[commandName]; exists {
				if command.Modal != nil {
					err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseModal,
						Data: &discordgo.InteractionResponseData{
							CustomID:   command.Modal.ID + "_" + i.Interaction.Member.User.ID,
							Title:      command.Modal.Title,
							Components: command.Modal.Components,
						},
					})
					if err != nil {
						log.Printf("Error showing modal: %v", err)
					}
				} else if command.Handler != nil {
					command.Handler(s, i)
				}
			}
		case discordgo.InteractionModalSubmit:
			data := i.ModalSubmitData()
			modalID := strings.Split(data.CustomID, "_")[0]

			for _, cmd := range types.RegisteredCommands {
				if cmd.Modal != nil && cmd.Modal.ID == modalID {
					cmd.Modal.Handler(s, i)
					return
				}
			}
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
