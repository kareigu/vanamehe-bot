package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var joinInfo = discordgo.ApplicationCommand{
	Name:        "join",
	Description: "Join your current voice channel",
}

func joinCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.GuildID == "" {
		InteractionRespondMessage(s, i, "mitte otsesõnumis")
		return
	}

	if s.VoiceConnections[i.GuildID] != nil {
		InteractionRespondMessage(s, i, "Juba kanalis")
		return
	}

	guild, err := s.State.Guild(i.GuildID)
	if err != nil {
		log.Printf("Error getting guildID: %v", err)
		InteractionRespondMessage(s, i, "Esines viga")
		return
	}

	for _, state := range guild.VoiceStates {
		if state.UserID == i.Member.User.ID {
			voice, err := s.ChannelVoiceJoin(guild.ID, state.ChannelID, false, false)
			if err != nil {
				InteractionRespondMessage(s, i, "Esines viga")
				log.Printf("Error joining voice: %v", err)
				return
			}
			ch, err := s.State.Channel(voice.ChannelID)
			if err != nil {
				InteractionRespondMessage(s, i, "Esines viga")
				log.Printf("Error getting voice channel: %v", err)
				return
			}

			err = InteractionRespondMessage(s, i, fmt.Sprintf("Liitunud kanaliga %v", ch.Name))
			if err != nil {
				InteractionRespondMessage(s, i, "Esines viga")
				log.Printf("Join issue: %v", err)
			}
			return
		}
	}
}
