package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {

	// parse the flags
	flags()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("Error creating Discord session: ", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	//SW Every time a message is received, it will call this function "messageHandler"
	dg.AddHandler(messageHandler)

	// Define the identity intent for this session
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()

	// defer the closing of the connection
	defer dg.Close()
	if err != nil {
		log.Println("Error opening connection: ", err)
		return
	}

	//SW This part just holds up the program until you try to terminate it,
	//     gracefully closes the Discord connection
	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")

	select {}

}
