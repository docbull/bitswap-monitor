package conn

import (
	"net/http"
	"os"
	"time"
)

type HttpClient struct {
	Client *http.Client
	URL    string
}

// func NewHTTPClient() fx.Option {
// 	return fx.Provide(func() *HttpClient {
// 		url := "http://localhost:5001"
// 		if len(os.Args) > 1 {
// 			url = os.Args[1]
// 		}
// 		return &HttpClient{
// 			Client: &http.Client{
// 				Transport:     nil,
// 				CheckRedirect: nil,
// 				Jar:           nil,
// 				Timeout:       time.Second * 10,
// 			},
// 			URL: url + "/api/v0/",
// 		}
// 	})
// }

func NewHTTPClient() *HttpClient {
	url := "http://localhost:5001"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}
	return &HttpClient{
		Client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       time.Second * 10,
		},
		URL: url + "/api/v0/",
	}
}
