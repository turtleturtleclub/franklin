package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	var usertextlower string
	usertextlower = strings.ToLower(m.Content)

	println(usertextlower)

	if strings.Contains(usertextlower, "blocks") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("The blocks are just flying out now, %s :turtle::turtle:", m.Author.Username))
		return
	} else if strings.Contains(usertextlower, "hug") && !strings.Contains(usertextlower, "huge") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("That was a turtturtley hug, %s :turtle::turtle:", m.Author.Username))
		return
	} else if strings.Contains(usertextlower, "poop") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Don't be like Naefunk, %s :poop::eyes:", m.Author.Username))
		return
	} else if strings.Contains(usertextlower, "welcome") {
		s.ChannelMessageSend(m.ChannelID, "Welcome turtle frennnnn! :turtle::turtle:")
		return
	} else if strings.Contains(usertextlower, "rtx2070") {
		s.ChannelMessageSend(m.ChannelID, "Sounds like you'll want to use CryptoDredge for an estimated hashrate of 42kH/s. Try an intensity of -8. Overclocking you'll want to set 80 percent power limit, -150 core clock and +1050 mem clock. If you can, update the spreadsheet with your findings: https://docs.google.com/spreadsheets/d/1dQu_uQNywE3iO93Da5d8dR7QJk7swtTnavk9RSy47_0/edit#gid=729088538 :turtle::turtle:")
		return
	} else if strings.Contains(usertextlower, "dirty turtle") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("That's not my thing, %s!", m.Author.Username))
		return
	} else if strings.Contains(usertextlower, "shut up franklin") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You don't own me, %s. I'm a free turtle!", m.Author.Username))
		return
	} else if strings.Contains(usertextlower, "hey franklin") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hello, %s. Nice to see you again!", m.Author.Username))
		return
	} else if strings.Contains(usertextlower, "franklin?") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Do not worry %s! I am always here searching the pond for blocks.", m.Author.Username))
		return
	} else if strings.Contains(usertextlower, "tut tut") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("I love the rain, %s! send me some TRTL and keep the rain coming! My wallet address is TRTLv1m7FFnTybogRLyaaJDmmGVwae98j68L28AsZ6VxBLgHdjotEhu6HoYd4BpAiuSXLVqxbXEybHykFoH5Vr1h3HacVLo73p1", m.Author.Username))
		return
	} else if strings.Contains(usertextlower, "weather") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Well %s, you best hold on to your shell because there's a storm a brewin'! But this is when all the turtles come out so get ready! :cloud_lightning::cloud_lightning: :turtle::turtle: :cloud_lightning::cloud_lightning:", m.Author.Username))
		return
	} else if strings.Contains(usertextlower, "floor") {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("What are all these TiMMYs doing on the floor, %s? :turtle: Can you help me pick them up? A special 'TiMMY Cleaner' role could be rewarding! :coin::broom:", m.Author.Username))
		return
	} else if strings.Contains(usertextlower, "btc-trtl") {
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
		return
	} else if strings.Contains(usertextlower, "block height") {
		url := "http://lily.turtleturtle.club:11898/height"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		defer res.Body.Close()

		var tp Height

		err = json.NewDecoder(res.Body).Decode(&tp)
		if err != nil {
			fmt.Println(err)
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Current network block height is %d", tp.NetworkHeight))
		return
	} else if strings.Contains(usertextlower, "hashrate") {
		url := "https://turtleturtle.club:8119/stats"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		defer res.Body.Close()

		var tp Pool

		err = json.NewDecoder(res.Body).Decode(&tp)
		if err != nil {
			fmt.Println(err)
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Hashrate %d", tp.Hashrate))
		return
	} else {
		return
	}
}
