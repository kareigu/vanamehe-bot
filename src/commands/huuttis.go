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
	err := InteractionRespondMessage(s, i, "https://mxrr.dev/files/vanamehe/huuttis.mp4")

	if err != nil {
		log.Printf("Huuttis issue: %v", err)
	}
}
