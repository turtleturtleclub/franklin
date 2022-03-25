package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Price struct {
	Success      bool   `json:"success,omitempty"`
	Initialprice string `json:"initialprice,omitempty"`
	Price        string `json:"price,omitempty"`
	High         string `json:"high,omitempty"`
	Low          string `json:"low,omitempty"`
	Volume       string `json:"volume,omitempty"`
	Bid          string `json:"bid,omitempty"`
	Ask          string `json:"ask,omitempty"`
}

// Variables used for command line parameters
var (
	Token string
)

func init() {
	//SW This part takes the command line arguments and combines them
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	//SW The program is called via run or ./ with the flag -t bot-token
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	//SW Every time a message is received, it will call this function "messageCreate"
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	//SW This part just holds up the program until you try to terminate it,
	//     gracefully closes the Discord connection
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
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

	var usertextlower string
	usertextlower = strings.ToLower(m.Content)

	println(usertextlower)

	if strings.Contains(usertextlower, "blocks") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The blocks are just flying out now, %s :turtle::turtle:", m.Author.Username))
	}
	if strings.Contains(usertextlower, "hug") && !strings.Contains(usertextlower, "huge") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("That was a turtturtley hug, %s :turtle::turtle:", m.Author.Username))
	}
	if strings.Contains(usertextlower, "poop") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Don't be like Naefunk, %s :poop::eyes:", m.Author.Username))
	}
	if strings.Contains(usertextlower, "welcome") {
		s.ChannelMessageSend(m.ChannelID, "Welcome turtle frennnnn! :turtle::turtle:")
	}
	if strings.Contains(usertextlower, "rtx2070") {
		s.ChannelMessageSend(m.ChannelID, "Sounds like you'll want to use CryptoDredge for an estimated hashrate of 42kH/s. Try an intensity of -8. Overclocking you'll want to set 80 percent power limit, -150 core clock and +1050 mem clock. If you can, update the spreadsheet with your findings: https://docs.google.com/spreadsheets/d/1dQu_uQNywE3iO93Da5d8dR7QJk7swtTnavk9RSy47_0/edit#gid=729088538 :turtle::turtle:")
	}
	if strings.Contains(usertextlower, "dirty turtle") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("That's not my thing, %s!", m.Author.Username))
	}
	if strings.Contains(usertextlower, "shut up franklin") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You don't own me, %s. I'm a free turtle!", m.Author.Username))
	}
	if strings.Contains(usertextlower, "hey franklin") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hello, %s. Nice to see you again!", m.Author.Username))
	}
	if strings.Contains(usertextlower, "franklin?") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Do not worry %s! I am always here searching the pond for blocks.", m.Author.Username))
	}
	if strings.Contains(usertextlower, "tut tut") {
                s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("I love the rain, %s! send me some TRTL and keep the rain coming! My wallet address is TRTLv1m7FFnTybogRLyaaJDmmGVwae98j68L28AsZ6VxBLgHdjotEhu6HoYd4BpAiuSXLVqxbXEybHykFoH5Vr1h3HacVLo73p1", m.Author.Username))
        }
	if strings.Contains(usertextlower, "weather") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Well %s, you best hold on to your shell because there's a storm a brewin'! But this is when all the turtles come out so get ready! :cloud_lightning::cloud_lightning: :turtle::turtle: :cloud_lightning::cloud_lightning:", m.Author.Username))
	}
  	if strings.Contains(usertextlower, "floor") {
                s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("What are all these TiMMYs doing on the floor, %s? :turtle: Can you help me pick them up? A special 'TiMMY Cleaner' role could be rewarding! :coin::broom:", m.Author.Username))
        }

	if strings.Contains(usertextlower, "btc-trtl") {
		url := "https://tradeogre.com/api/v1/ticker/BTC-TRTL"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		defer res.Body.Close()

		var tp Price

		err = json.NewDecoder(res.Body).Decode(&tp)
		if err != nil {
			fmt.Println(err)
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("1 :turtle: to BTC is %s ", tp.Ask))
	}
}
