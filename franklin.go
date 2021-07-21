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

	if strings.Contains(usertextlower, "block") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Patience is a virtue, %s :turtle::turtle:", m.Author.Username))
	}
	if strings.Contains(usertextlower, "hug") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("That was a turtturtley hug, %s :turtle::turtle:", m.Author.Username))
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
