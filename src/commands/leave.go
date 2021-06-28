package commands

import (
	"github.com/bwmarrin/discordgo"
)

var leaveInfo = discordgo.ApplicationCommand{
	Name:        "leave",
	Description: "Leave the voice channel",
}

func leaveCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.GuildID == "" {
		InteractionRespondMessage(s, i, "mitte otsesõnumis")
		return
	}

	if s.VoiceConnections[i.GuildID] != nil {
		s.VoiceConnections[i.GuildID].Disconnect()
		InteractionRespondMessage(s, i, "Vasak kanal")
	} else {
		InteractionRespondMessage(s, i, "Mitte kanalis")
	}
}
