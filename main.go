package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	bitswap_stat "github.com/docbull/bitswap-monitor/bitswap-stat"
	"github.com/docbull/bitswap-monitor/conn"
	"github.com/rivo/tview"
)

type BitswapStat bitswap_stat.BitswapStat

// RefreshMonitor re-rendering bitswap logs every 10ms.
func RefreshMonitor(client *conn.HttpClient, bitswapStat *BitswapStat, view *tview.TextView) {
	text := ""
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

	for i := 0; i < len(bitswapStat.Wantlist); i++ {
		text += fmt.Sprintf("%v ", bitswapStat.Wantlist[i])
	}
	// if len(bitswapStat.Wantlist) == 0 {
	// 	text += fmt.Sprintf("%s", time.Now())
	// }
	fmt.Fprintf(view, "%s\n", text)
}

func main() {
	client := conn.NewHTTPClient()
	var bitswapStat *BitswapStat

	app := tview.NewApplication()
	view := tview.NewTextView()
	view.SetBorder(true)
	view.SetChangedFunc(func() {
		app.Draw()
	})
	view.SetScrollable(true)

	go func() {
		for {
			RefreshMonitor(client, bitswapStat, view)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	if err := app.SetRoot(view, true).SetFocus(view).Run(); err != nil {
		panic(err)
	}
}
