package utils

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

const PRIIDIK_ID = "856586054048022549"

var VOICE_LINES = [10]string{"se01",
	"se02", "se03", "se04",
	"se05", "se06", "se07",
	"se08", "se09", "se10"}

func PlayVoiceLine(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, conn := range s.VoiceConnections {
		guild, err := s.State.Guild(conn.GuildID)
		if err != nil {
			log.Print("Error getting guild")
			return
		}

		for _, vs := range guild.VoiceStates {
			if vs.UserID == PRIIDIK_ID {
				rand.Seed(time.Now().UnixNano())
				roll := rand.Intn(10)
				log.Printf("Voiceline roll: %v", roll)
				filename := fmt.Sprintf("./assets/%v.mp3", VOICE_LINES[roll])
				dgvoice.PlayAudioFile(conn, filename, make(chan bool))
				log.Print("Ran se on")
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Playing %v", filename))
			}
		}
	}
}
