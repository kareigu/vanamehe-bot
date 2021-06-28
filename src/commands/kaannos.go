package commands

import (
	"log"
	"utils"

	"github.com/bwmarrin/discordgo"
)

var kaannosInfo = discordgo.ApplicationCommand{
	Name:        "käännös",
	Description: "Käännös mint",
}

func kaannosCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := InteractionRespondMessage(s, i, utils.Videos[utils.KÄÄNNÖS])

	if err != nil {
		log.Printf("Käännös issue: %v", err)
	}
}
