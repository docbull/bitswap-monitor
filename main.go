package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
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
		panic(err)
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

	// fmt.Fprintf(view, "%s ", time.Now())
	for i := 0; i < len(bitswapStat.Wantlist); i++ {
		a := reflect.ValueOf(bitswapStat.Wantlist[i])
		for _, b := range a.MapKeys() {
			c := a.MapIndex(b)
			t := c.Interface()
			text += fmt.Sprintf("%v ", t)
		}
	}
	if text != "" {
		fmt.Fprintf(view, "%s\n\n", text)
	}
}

func main() {
	client := conn.NewHTTPClient()
	var bitswapStat *BitswapStat

	app := tview.NewApplication()
	newPrimitive := func(text string) *tview.TextView {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text).
			SetChangedFunc(func() { app.Draw() })
	}
	// menu := newPrimitive("Menu")
	info := newPrimitive("IPFS Bitswap Monitor \n\n")
	main := newPrimitive("Want List \n\n")
	// sideBar := newPrimitive("Side Bar")

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(info, 0, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(main, 1, 0, 1, 3, 0, 0, false)

	go func() {
		for {
			RefreshMonitor(client, bitswapStat, main)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Layout for screens wider than 100 cells.
	// grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
	// 	AddItem(main, 1, 1, 1, 1, 0, 100, false).
	// 	AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		fmt.Printf("\n %v\n", err)
		os.Exit(1)
	}
}
