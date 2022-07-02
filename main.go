package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/Sab94/ipfs-monitor/types"
	"github.com/Sab94/ipfs-monitor/widget"
)

type BitswapStatBlock widget.Widget

type HttpClient struct {
	Client *http.Client
	Base   string
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
	fmt.Println("client:", client)

	// text := ""
	var data = []byte{}
	var bitswapStat types.BitswapStat

	req, err := http.NewRequest("POST", client.Base+"bitswap/stat", nil)
	resp, err := client.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	data, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &bitswapStat)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("%12s: [green]%-7d[white]%12s: [green]%-7d[white]%12s: [green]%-7d\n",
		"Blocks Got", bitswapStat.BlocksReceived, "Blocks Sent",
		bitswapStat.BlocksSent, "Dup Blocks", bitswapStat.DupBlksReceived)
}
