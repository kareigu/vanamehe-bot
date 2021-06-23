package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var jattisInfo = discordgo.ApplicationCommand{
	Name:        "jättis",
	Description: "Jättis huuttis mint",
}

func jattisCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := InteractionRespondMessage(s, i, "https://mxrr.dev/files/vanamehe/j%C3%A4ttis.mp4")

	if err != nil {
		log.Printf("Jättis issue: %v", err)
	}
}
