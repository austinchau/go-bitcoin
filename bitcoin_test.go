package bitcoin

import (
  "testing"
)

var (
  username string
  password string
  port int
  host string
)

func init() {
  username = "" // rpcuser in your bitcoin.conf
  password = "" // rpcpassword in your bitcoin.conf
  port = 18332 // default to testnet port.  Change to 8332 for live blockchain
  host = "localhost"
}

func TestNewClient(t *testing.T) {
  c := NewClient(username, password, host, port)
  if c == nil {
    t.Error()
  }
}

func TestMakeRequest(t *testing.T) {
  c := NewClient(username, password, host, port)
  if c == nil {
    t.Error()
  }

  params := []string{}
  _, err := c.MakeRequest("getinfo", params)
  if err != nil {
    t.Error(err)
  }
}