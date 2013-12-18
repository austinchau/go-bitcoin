package bitcoin

import (
  "testing"
  "fmt"
)

func TestNewClient(t *testing.T) {
  c := NewClient("", "", "localhost", 18332)
  if c == nil {
    t.Error()
  }
}

func TestMakeRequest(t *testing.T) {
  c := NewClient("", "", "localhost", 18332)
  if c == nil {
    t.Error()
  }

  params := []string{}
  result, err := c.MakeRequest("getinfo", params)
  if err != nil {
    t.Error(err)
  }
  fmt.Println(result)
}