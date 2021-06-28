package commands

import (
	"log"
	"utils"

	"github.com/bwmarrin/discordgo"
)

var kiinnostaaInfo = discordgo.ApplicationCommand{
	Name:        "kiinnostaa",
	Description: "Kiinnostaa, ei kiinnosta",
}

func kiinnostaaCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := InteractionRespondMessage(s, i, utils.Videos[utils.KIINNOSTAA])

	if err != nil {
		log.Printf("Kiinnostaa issue: %v", err)
	}
}
