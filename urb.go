package main

import (
  "fmt"
  "os"
  "net/http"
  "encoding/json"
  "github.com/urfave/cli"
)

type udResponse struct {
  List []udDefinition `json: list`
}

type udDefinition struct {
  Definition string `json: definition`
}

func main() {
  const urbandictionaryApi = "http://api.urbandictionary.com/v0/define?term="
  app := cli.NewApp()
  app.Name = "urb"
  app.Usage = "lookup urbandictionary.com from the command line"
  app.Action = func(c *cli.Context) error {
    str := c.Args().Get(0)

    resp, err := http.Get(urbandictionaryApi + str)
    defer resp.Body.Close()
    if err != nil {
    	return nil
    }

    var f udResponse
    err = json.NewDecoder(resp.Body).Decode(&f)
    if err != nil {
      return nil
    }

    fmt.Println("Looking up: " + str)
    fmt.Println(f.List[0].Definition)

    return nil
  }

  app.Run(os.Args)
}
