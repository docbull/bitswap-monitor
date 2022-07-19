package bitswap_stat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/docbull/bitswap-monitor/conn"
	"github.com/rivo/tview"
)

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

// RefreshMonitor re-rendering bitswap logs every 500ms.
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
