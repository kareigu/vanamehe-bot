package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var kiinnostaaInfo = discordgo.ApplicationCommand{
	Name:        "kiinnostaa",
	Description: "Kiinnostaa, ei kiinnosta",
}

func kiinnostaaCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: "https://mxrr.dev/files/vanamehe/kiinnostaa.mp4",
		},
	})

	if err != nil {
		log.Printf("Kiinnostaa issue: %v", err)
	}
}
