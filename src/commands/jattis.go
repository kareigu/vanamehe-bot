package commands

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var jattisInfo = discordgo.ApplicationCommand{
	Name:        "jättis",
	Description: "Jättis huuttis mint",
}

func jattisCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel := i.ChannelID

	file, err := os.Open("assets/jättis.mp4")
	if err != nil {
		return
	}

	v := discordgo.File{
		Name:        "jättis.mp4",
		ContentType: "video/mp4",
		Reader:      file,
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionApplicationCommandResponseData{
			Content: "jättis huuttis mint",
		},
	})

	if err != nil {
		log.Printf("Jättis issue: %v", err)
	}

	s.ChannelMessageSendComplex(channel, &discordgo.MessageSend{
		Files: []*discordgo.File{
			&v,
		},
	})
}
