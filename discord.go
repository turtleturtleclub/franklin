package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID != s.State.User.ID {

		usertextlower := strings.ToLower(m.Content)

		parseMessage(s, m, usertextlower)

	} else {
		log.Println("Ignored input: " + m.Content)
		return
	}
}

// This function takes inbound messages and responds to known patterns with a programmed response
func parseMessage(s *discordgo.Session, m *discordgo.MessageCreate, messageInput string) {

	if strings.Contains(messageInput, "blocks") {
		log.Println("Received \"blocks\" command")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The blocks are just flying out now, %s :turtle::turtle:", m.Author.Username))
		return
	} else if strings.Contains(messageInput, "hug") && !strings.Contains(messageInput, "huge") {
		log.Println("Received \"hug\" command")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("That was a turtturtley hug, %s :turtle::turtle:", m.Author.Username))
		return
	} else if strings.Contains(messageInput, "poop") {
		log.Println("Received \"poop\" command")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Don't be like Naefunk, %s :poop::eyes:", m.Author.Username))
		return
	} else if strings.Contains(messageInput, "welcome") {
		log.Println("Received \"welcome\" command")
		s.ChannelMessageSend(m.ChannelID, "Welcome turtle frennnnn! :turtle::turtle:")
		return
	} else if strings.Contains(messageInput, "rtx2070") {
		log.Println("Received \"rtx2070\" command")
		s.ChannelMessageSend(m.ChannelID, "Sounds like you'll want to use CryptoDredge for an estimated hashrate of 42kH/s. Try an intensity of -8. Overclocking you'll want to set 80 percent power limit, -150 core clock and +1050 mem clock. If you can, update the spreadsheet with your findings: https://docs.google.com/spreadsheets/d/1dQu_uQNywE3iO93Da5d8dR7QJk7swtTnavk9RSy47_0/edit#gid=729088538 :turtle::turtle:")
		return
	} else if strings.Contains(messageInput, "dirty turtle") {
		log.Println("Received \"dirty turtle\" command")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("That's not my thing, %s!", m.Author.Username))
		return
	} else if strings.Contains(messageInput, "shut up franklin") {
		log.Println("Received \"shut up franklin\" command")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You don't own me, %s. I'm a free turtle!", m.Author.Username))
		return
	} else if strings.Contains(messageInput, "hey franklin") {
		log.Println("Received \"hey franklin\" command")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hello, %s. Nice to see you again!", m.Author.Username))
		return
	} else if strings.Contains(messageInput, "franklin?") {
		log.Println("Received \"franklin?\" command")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Do not worry %s! I am always here searching the pond for blocks.", m.Author.Username))
		return
	} else if strings.Contains(messageInput, "tut tut") {
		log.Println("Received \"tut tut\" command")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("I love the rain, %s! send me some TRTL and keep the rain coming! My wallet address is TRTLv1m7FFnTybogRLyaaJDmmGVwae98j68L28AsZ6VxBLgHdjotEhu6HoYd4BpAiuSXLVqxbXEybHykFoH5Vr1h3HacVLo73p1", m.Author.Username))
		return
	} else if strings.Contains(messageInput, "weather") {
		log.Println("Received \"weather\" command")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Well %s, you best hold on to your shell because there's a storm a brewin'! But this is when all the turtles come out so get ready! :cloud_lightning::cloud_lightning: :turtle::turtle: :cloud_lightning::cloud_lightning:", m.Author.Username))
		return
	} else if strings.Contains(messageInput, "floor") {
		log.Println("Received \"floor\" command")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("What are all these TiMMYs doing on the floor, %s? :turtle: Can you help me pick them up? A special 'TiMMY Cleaner' role could be rewarding! :coin::broom:", m.Author.Username))
		return
	} else if strings.Contains(messageInput, "btc-trtl") {
		log.Println("Received \"btc-trtl\" command")

		httpGetResponse := httpGet("https://tradeogre.com/api/v1/ticker/BTC-TRTL")

		defer httpGetResponse.Body.Close()

		var tp Price

		json.NewDecoder(httpGetResponse.Body).Decode(&tp)

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("1 :turtle: to BTC is %s ", tp.Ask))

		return
	} else if strings.Contains(messageInput, "block height") {
		log.Println("Received \"block height\" command")

		httpGetResponse := httpGet("http://lily.turtleturtle.club:11898/height")

		defer httpGetResponse.Body.Close()

		var height Height

		json.NewDecoder(httpGetResponse.Body).Decode(&height)

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Current network block height is %d", height.NetworkHeight))

		return
	} else if strings.Contains(messageInput, "hashrate") {
		log.Println("Received \"hashrate\" command")

		httpGetResponse := httpGet("https://turtleturtle.club:8119/stats")

		defer httpGetResponse.Body.Close()

		var tp Pool

		json.NewDecoder(httpGetResponse.Body).Decode(&tp)

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hashrate %d", tp.Hashrate))
		return
	}
}
