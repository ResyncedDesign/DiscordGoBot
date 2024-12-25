package modals

import (
	"DiscordGoBot/src/types"
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	permission := int64(discordgo.PermissionSendMessages)

	surveyModal := &types.Modal{
		ID:    "survey_modal",
		Title: "Feedback Survey",
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:  "feedback",
						Label:     "Your Feedback",
						Style:     discordgo.TextInputParagraph,
						Required:  true,
						MaxLength: 1000,
					},
				},
			},
		},
		Handler: handleSurveySubmit,
	}

	types.RegisteredCommands["survey"] = &types.Command{
		Name:                     "survey",
		Description:              "Open feedback survey",
		Category:                 "modals",
		Modal:                    surveyModal,
		DefaultMemberPermissions: &permission,
		Global:                   true,
	}
}

func handleSurveySubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()
	feedback := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	log.Printf("Feedback received: %s", feedback)

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Thank you for your feedback!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		log.Printf("Error responding to modal: %v", err)
	}
}
