package main

// Price is a struct model of the JSON response from the price API from tradeogre
type Price struct {
	Success      bool   `json:"success,omitempty"  db:"success"`
	Initialprice string `json:"initialprice,omitempty"  db:"initialprice"`
	Price        string `json:"price,omitempty"  db:"price"`
	High         string `json:"high,omitempty"  db:"high"`
	Low          string `json:"low,omitempty"  db:"low"`
	Volume       string `json:"volume,omitempty"  db:"volume"`
	Bid          string `json:"bid,omitempty"  db:"bid"`
	Ask          string `json:"ask,omitempty"  db:"ask"`
}

// Height is the current block height of the network, as supplied by the daemon
type Height struct {
	NetworkHeight int32 `json:"networkHeight,omitempty"  db:"network_height"`
}

// Pool is the hashrate and last block response from the pool daemon
type Pool struct {
	Hashrate  int32 `json:"hashrate,omitempty"  db:"hashrate"`
	LastBlock int64 `json:"lastBlockFoundprop,omitempty"  db:"last_block"`
}
