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

type Height struct {
	NetworkHeight int32 `json:"networkHeight,omitempty"`
}

type Pool struct {
	Hashrate  int32 `json:"hashrate,omitempty"`
	LastBlock int64 `json:"lastBlockFoundprop,omitempty"`
}

// Variables used for command line parameters
var (
	Token string
)
