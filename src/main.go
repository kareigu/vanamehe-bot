package main

import (
	"commands"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	client  *discordgo.Session
	GuildID string
)

const PRIIDIK_ID = "856586054048022549"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load env file")
	}

	bot_token := os.Getenv("TOKEN")
	GuildID = os.Getenv("GUILD_ID")

	if bot_token == "" {
		log.Fatal("Bot token missing from .env")
	}

	client, err = discordgo.New("Bot " + bot_token)
	if err != nil {
		log.Fatal("Couldn't create bot client")
	}
}

func init() {
	client.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.Handlers[i.Data.Name]; ok {
			h(s, i)
			log.Printf("%v#%v ran command %v",
				i.Member.User.Username,
				i.Member.User.ID,
				i.Data.Name)
		}
	})

	client.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Author.ID == PRIIDIK_ID && m.GuildID == "" {
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
	})

	//m := discordgo.DM
}

func main() {
	client.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuilds | discordgo.IntentsGuildVoiceStates | discordgo.IntentsDirectMessages
	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Running...")
	})

	err := client.Open()
	if err != nil {
		log.Fatal("Couldn't start bot client")
	}

	for _, v := range commands.List {
		_, err = client.ApplicationCommandCreate(client.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Couldn't create command %v : %v", v.Name, err)
		}
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}
