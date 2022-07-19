package peerinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/docbull/bitswap-monitor/conn"
	"github.com/rivo/tview"
)

type PeerInfo struct {
	Addresses       []string `json:"Addresses"`
	AgentVersion    string   `json:"AgentVersion"`
	ID              string   `json:"ID"`
	ProtocolVersion string   `json:"ProtocolVersion"`
	PublicKey       string   `json:"PublicKey"`
}

// ShowPeerInfo prints IPFS node's ID
func ShowPeerInfo(client *conn.HttpClient, peerInfo *PeerInfo, view *tview.TextView) {
	text := ""
	var data = []byte{}
	req, err := http.NewRequest("POST", client.URL+"id", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	data, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &peerInfo)
	if err != nil {
		fmt.Println(err)
	}

	text += fmt.Sprintf("ID: %v", peerInfo.ID)
	fmt.Fprintf(view, "%s", text)
}
