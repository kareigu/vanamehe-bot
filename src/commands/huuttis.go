package commands

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var huuttisInfo = discordgo.ApplicationCommand{
	Name:        "huuttis",
	Description: "Huuttis mint",
}

func huuttisCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel := i.ChannelID

	file, err := os.Open("assets/huuttis.mp4")
	if err != nil {
		return
	}

	v := discordgo.File{
		Name:        "huuttis.mp4",
		ContentType: "video/mp4",
		Reader:      file,
	}

	s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Files: []*discordgo.File{
			&v,
		},
	})

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: "huuttis mint",
		},
	})

	if err != nil {
		log.Printf("Huuttis issue: %v", err)
	}
}
