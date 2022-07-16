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
		for _, v := range bitswapStat.Wantlist {
			a := reflect.ValueOf(v)
			for _, b := range a.MapKeys() {
				c := a.MapIndex(b)
				t := c.Interface()
				text += fmt.Sprintf("%v ", t)
			}
			// val := reflect.ValueOf(v).Elem()
			// n := val.FieldByName("/").Interface().(string)
			// text += fmt.Sprintf("%v ", n)
			// text += fmt.Sprintf("%v ", b)
		}
		// text += fmt.Sprintf("%v ", bitswapStat.Wantlist[i])
	}
	if text != "" {
		fmt.Fprintf(view, "%s\n\n", text)
	}
}

func main() {
	client := conn.NewHTTPClient()
	var bitswapStat *BitswapStat

	// app := tview.NewApplication()
	// view := tview.NewTextView()
	// view.SetBorder(true)
	// view.SetChangedFunc(func() {
	// 	app.Draw()
	// })
	// view.SetScrollable(true)

	// go func() {
	// 	for {
	// 		RefreshMonitor(client, bitswapStat, view)
	// 		time.Sleep(500 * time.Millisecond)
	// 	}
	// }()

	// if err := app.SetRoot(view, true).SetFocus(view).Run(); err != nil {
	// 	panic(err)
	// }

	app := tview.NewApplication()
	newPrimitive := func(text string) *tview.TextView {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text).
			SetChangedFunc(func() { app.Draw() })
	}
	// menu := newPrimitive("Menu")
	main := newPrimitive("Want List")
	// sideBar := newPrimitive("Side Bar")

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("IPFS Bitswap Monitor"), 0, 0, 1, 3, 0, 0, false)

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
