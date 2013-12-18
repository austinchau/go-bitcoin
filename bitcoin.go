package bitcoin

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "time"
  "strings"
  "errors"
  "encoding/json"
)

type JsonRpc struct {
  User string
  Password string
  Host string
  Port int
}

func (c *JsonRpc) MakeRequest(method string, params []string) (map[string]interface{}, error) {
  baseUrl := fmt.Sprintf("http://%s:%d", c.Host, c.Port)
  client := new(http.Client)
  req, err := http.NewRequest("POST", baseUrl, nil)
  if err != nil {
    return nil, err
  }
  
  req.SetBasicAuth(c.User, c.Password)
  req.Header.Add("Content-Type", "text/plain")
  
  args := make(map[string]interface{})
  args["jsonrpc"] = "1.0"
  args["id"] = time.Now().UnixNano()
  args["method"] = method
  args["params"] = params
  
  j, err := json.Marshal(args)
  if err != nil {
    fmt.Println(err)
  }
  
  req.Body = ioutil.NopCloser(strings.NewReader(string(j)))
  req.ContentLength = int64(len(string(j)))
  
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  
  defer resp.Body.Close()
  bytes, _ := ioutil.ReadAll(resp.Body)
  
  var data map[string]interface{}
  json.Unmarshal(bytes, &data)
  if err, found := data["error"]; found && err != nil {
    str,_ := json.Marshal(err)
    return nil, errors.New(string(str))
  }
  
  if result, found := data["result"]; found {
    return result.(map[string]interface{}), nil
  } else {
    return nil, errors.New("no result")
  }
}

func NewClient(user string, password string, host string, port int) *JsonRpc {
  c := JsonRpc{user, password, host, port}
  return &c
}