package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sab94/ipfs-monitor/widget"
	"github.com/docbull/bitswap-monitor/conn"
)

type BitswapStatBlock widget.Widget

type BitswapStat struct {
	BlocksReceived   uint64        `json:"BlocksReceived"`
	BlocksSent       uint64        `json:"BlocksSent"`
	DataReceived     uint64        `json:"DataReceived"`
	DataSent         uint64        `json:"DataSent"`
	DupBlksReceived  uint64        `json:"DupBlksReceived"`
	DupDataReceived  uint64        `json:"DupDataReceived"`
	MessagesReceived uint64        `json:"MessagesReceived"`
	Peers            []interface{} `json:"Peers"`
	ProvideBufLen    int           `json:"ProvideBufLen"`
	Wantlist         []interface{} `json:"Wantlist"`
}

func main() {
	fmt.Println("Bitswap Monitor ...")

	client := conn.NewHTTPClient()
	// fmt.Println("client:", client)

	// text := ""
	var data = []byte{}
	var bitswapStat *BitswapStat

	req, err := http.NewRequest("POST", client.URL+"bitswap/stat", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	data, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &bitswapStat)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Blocks Got:", bitswapStat.BlocksReceived)
	fmt.Println("Blocks Sent:", bitswapStat.BlocksSent)
	fmt.Println("Duplicated Block Received:", bitswapStat.DupBlksReceived)

	fmt.Println("Data Received:", bitswapStat.DataReceived)
	fmt.Println("Data Sent:", bitswapStat.DataSent)
	fmt.Println("Duplicated Data Received:", bitswapStat.DupDataReceived)

	fmt.Println("Messages Received:", bitswapStat.MessagesReceived)
	// fmt.Println("Peers:", bitswapStat.Peers)
	fmt.Println("Provide Buffer Length:", bitswapStat.ProvideBufLen)
	fmt.Println("Wantlist:", bitswapStat.Wantlist)
}
