package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var homoInfo = discordgo.ApplicationCommand{
	Name:        "homo",
	Description: "jätti homo mint",
}

func homoCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := InteractionRespondMessage(s, i, "https://mxrr.dev/files/vanamehe/homo.mp4")

	if err != nil {
		log.Printf("Homo issue: %v", err)
	}
}
