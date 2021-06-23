package utils

import (
	"log"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

const PRIIDIK_ID = "856586054048022549"

func PlayVoiceLine(s *discordgo.Session) {
	for _, conn := range s.VoiceConnections {
		guild, err := s.State.Guild(conn.GuildID)
		if err != nil {
			log.Print("Error getting guild")
			return
		}

		for _, vs := range guild.VoiceStates {
			if vs.UserID == PRIIDIK_ID {
				dgvoice.PlayAudioFile(conn, "./assets/se01.mp3", make(chan bool))
				log.Print("Ran se on")
			}
		}
	}
}
