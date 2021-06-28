package commands

import (
	"log"
	"utils"

	"github.com/bwmarrin/discordgo"
)

var huuttisInfo = discordgo.ApplicationCommand{
	Name:        "huuttis",
	Description: "Huuttis mint",
}

func huuttisCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := InteractionRespondMessage(s, i, utils.Videos[utils.HUUTTIS])

	if err != nil {
		log.Printf("Huuttis issue: %v", err)
	}
}
