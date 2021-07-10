package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

func main() {

	fmt.Println("Enter username: ")

	var username string
	fmt.Scanln(&username)

	fmt.Println("Enter statement: ")

	var usertext string
	fmt.Scanln(&usertext)

	var usertextlower string
	usertextlower = strings.ToLower(usertext)

	if strings.Contains(usertextlower, "block") {
		fmt.Println("Patience is a virtue", username)
	}
	if strings.Contains(usertextlower, "hug") {
		fmt.Println("That was a turtturtley hug", username)
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

		fmt.Println(tp.Ask)
		fmt.Println(tp)
	}
}
