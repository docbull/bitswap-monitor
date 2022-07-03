package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Sab94/ipfs-monitor/widget"
)

type BitswapStatBlock widget.Widget

type HttpClient struct {
	Client *http.Client
	Base   string
}

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

func NewHTTPClient() *HttpClient {
	base := "http://localhost:5001"
	if len(os.Args) > 1 {
		base = os.Args[1]
	}
	return &HttpClient{
		Client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       time.Second * 10,
		},
		Base: base + "/api/v0/",
	}
}

func main() {
	fmt.Println("Bitswap Monitor ...")

	client := NewHTTPClient()
	// fmt.Println("client:", client)

	// text := ""
	var data = []byte{}
	var bitswapStat *BitswapStat

	req, err := http.NewRequest("POST", client.Base+"bitswap/stat", nil)
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

	// text += fmt.Sprintf("%12s: [green]%-7d[white]%12s: [green]%-7d[white]%12s: [green]%-7d\n",
	// 	"Blocks Got", bitswapStat.BlocksReceived, "Blocks Sent",
	// 	bitswapStat.BlocksSent, "Dup Blocks", bitswapStat.DupBlksReceived)
}
