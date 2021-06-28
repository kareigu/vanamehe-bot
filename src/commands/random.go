package commands

import (
	"log"
	"math/rand"
	"time"
	"utils"

	"github.com/bwmarrin/discordgo"
)

var randomInfo = discordgo.ApplicationCommand{
	Name:        "random",
	Description: "Random huuttis",
}

func randomCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(utils.VideoCount)

	err := InteractionRespondMessage(s, i, utils.Videos[roll])

	if err != nil {
		log.Printf("Random issue: %v", err)
	}
}
