package commands

import (
	"github.com/bwmarrin/discordgo"
)

var huuttisInfo = discordgo.ApplicationCommand{
	Name:        "huuttis",
	Description: "Huuttis mint",
}

func huuttisCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: "huuttis mint",
		},
	})
}
