package commands

import (
	"log"
	"utils"

	"github.com/bwmarrin/discordgo"
)

var homoInfo = discordgo.ApplicationCommand{
	Name:        "homo",
	Description: "jätti homo mint",
}

func homoCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := InteractionRespondMessage(s, i, utils.Videos[utils.HOMO])

	if err != nil {
		log.Printf("Homo issue: %v", err)
	}
}
