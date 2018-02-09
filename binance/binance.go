/*
   binance.go
       Wrapper for the Binance Exchange API

   Authors:
       Pat DePippo  <patrick.depippo@dcrypt.io>
       Matthew Woop <matthew.woop@dcrypt.io>

   To Do:

*/
package binance

//"errors"

const (
	// BaseURL for binance
	BaseURL = "https://api.binance.com"
)

// Binance stores the Binance Client
type Binance struct {
	client *Client
}

// New returns the Binance Client
func New(key, secret string) *Binance {
	client := NewClient(key, secret)
	return &Binance{client}
}
