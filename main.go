package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Sab94/ipfs-monitor/widget"
	bitswap_stat "github.com/docbull/bitswap-monitor/bitswap-stat"
	"github.com/docbull/bitswap-monitor/conn"
	"github.com/rivo/tview"
)

type BitswapStat bitswap_stat.BitswapStat

type BitswapStatBlock widget.Widget

// RefreshMonitor re-rendering bitswap logs every 10ms.
func RefreshMonitor(client *conn.HttpClient, bitswapStat *BitswapStat, textView *tview.TextView) {
	text := ""
	// fmt.Fprintf(textView, "%s", "ABSc")

	var data = []byte{}

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

	// var str []string
	for i := 0; i < len(bitswapStat.Wantlist); i++ {
		str := fmt.Sprintf("%v", bitswapStat.Wantlist[i])
		text += str
	}
	fmt.Fprintf(textView, "%s\n", text)
	// text += fmt.Sprintf("%s", text)
	// fmt.Println("Blocks Got:", bitswapStat.BlocksReceived)
	// fmt.Println("Blocks Sent:", bitswapStat.BlocksSent)
	// fmt.Println("Duplicated Block Received:", bitswapStat.DupBlksReceived)

	// fmt.Println("Data Received:", bitswapStat.DataReceived)
	// fmt.Println("Data Sent:", bitswapStat.DataSent)
	// fmt.Println("Duplicated Data Received:", bitswapStat.DupDataReceived)

	// fmt.Println("Messages Received:", bitswapStat.MessagesReceived)
	// // fmt.Println("Peers:", bitswapStat.Peers)
	// fmt.Println("Provide Buffer Length:", bitswapStat.ProvideBufLen)
	// fmt.Println("Wantlist:", bitswapStat.Wantlist)
	// fmt.Println("----------------------------------------------")
}

func main() {
	client := conn.NewHTTPClient()
	var bitswapStat *BitswapStat

	app := tview.NewApplication()
	textView := tview.NewTextView()
	textView.SetBorder(true)
	go func() {
		for {
			RefreshMonitor(client, bitswapStat, textView)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
		panic(err)
	}

	// for {
	// 	RefreshMonitor(client, bitswapStat)
	// 	time.Sleep(time.Millisecond * 1000)
	// }
}
