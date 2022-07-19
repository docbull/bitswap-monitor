package main

import (
	"fmt"
	"os"
	"time"

	bitswap_stat "github.com/docbull/bitswap-monitor/bitswap-stat"
	conn "github.com/docbull/bitswap-monitor/conn"
	peerinfo "github.com/docbull/bitswap-monitor/peer-info"
	"github.com/rivo/tview"
)

type BitswapStat *bitswap_stat.BitswapStat
type PeerInfo *peerinfo.PeerInfo

func main() {
	client := conn.NewHTTPClient()
	var bitswapStat BitswapStat
	var peerInfo PeerInfo

	app := tview.NewApplication()
	newPrimitive := func(text string) *tview.TextView {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text).
			SetChangedFunc(func() { app.Draw() })
	}

	info := newPrimitive("IPFS Bitswap Monitor \n\n")
	peerinfo.ShowPeerInfo(client, peerInfo, info)
	main := newPrimitive("Want List \n\n")

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(info, 0, 0, 1, 3, 0, 0, false)

	grid.AddItem(main, 1, 0, 1, 3, 0, 0, false)

	go func() {
		for {
			bitswap_stat.RefreshMonitor(client, bitswapStat, main)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		fmt.Printf("\n %v\n", err)
		os.Exit(1)
	}
}
