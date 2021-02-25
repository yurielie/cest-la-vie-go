package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

func init(){
	flag.StringVar(&envpath, "env", ".env", ".env file written the Bot Token")
	flag.Parse()
}

var (
	envpath string
)

func main(){

	err:= godotenv.Load(envpath)
	if err != nil {
		log.Fatal("Error loading .env file!")
		return
	}
	token := os.Getenv("DISCORD_TOKEN")

	dg, err := discordgo.New("Bot "+ token)
	if err != nil {
		fmt.Println("error crerating  Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	// Server名を取得して返します。
	if m.Content == "ServerName" {
		g, err := s.Guild(m.GuildID)
		if err != nil {
				log.Fatal(err)
		}
		log.Println(g.Name)
		s.ChannelMessageSend(m.ChannelID, g.Name)
	}

	// !Helloというチャットがきたら　「Hello」　と返します
	if m.Content == "!Hello" {
		s.ChannelMessageSend(m.ChannelID, "Hello")
	}
}
