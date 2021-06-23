package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var huuttisInfo = discordgo.ApplicationCommand{
	Name:        "huuttis",
	Description: "Huuttis mint",
}

func huuttisCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: "https://mxrr.dev/files/vanamehe/huuttis.mp4",
		},
	})

	if err != nil {
		log.Printf("Huuttis issue: %v", err)
	}
}
