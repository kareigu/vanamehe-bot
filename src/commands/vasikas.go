package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var vasikasInfo = discordgo.ApplicationCommand{
	Name:        "vasikas",
	Description: "vasikas",
}

func vasikasCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := InteractionRespondMessage(s, i, "https://www.youtube.com/watch?v=5RscteTVFtM")

	if err != nil {
		log.Printf("Vasikas issue: %v", err)
	}
}
