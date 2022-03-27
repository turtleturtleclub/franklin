package main

import "flag"

// Variables used for command line parameters
var (
	Token string
)

// This function parses the run flags for this session
func flags() {
	//SW This part takes the command line arguments and combines them
	flag.StringVar(&Token, "t", "", "Bot Token - This is the discord bot token that allows this program to speak in your channel as a bot.")
	flag.Parse()
	//SW The program is called via run or ./ with the flag -t bot-token
}
