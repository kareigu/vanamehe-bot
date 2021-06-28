package commands

import (
	"log"
	"utils"

	"github.com/bwmarrin/discordgo"
)

var jattisInfo = discordgo.ApplicationCommand{
	Name:        "jättis",
	Description: "Jättis huuttis mint",
}

func jattisCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := InteractionRespondMessage(s, i, utils.Videos[utils.JÄTTIS])

	if err != nil {
		log.Printf("Jättis issue: %v", err)
	}
}
