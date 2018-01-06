package main

import (
  "fmt"
  "os"
  "net/http"
  "encoding/json"
  "github.com/urfave/cli"
  "net/url"
)

type udResponse struct {
  List []udDefinition `json: list`
}

type udDefinition struct {
  Definition string `json: definition`
  Example string `json: example`
}

func main() {
  const urbandictionaryApi = "http://api.urbandictionary.com/v0/define?term="
  app := cli.NewApp()
  app.Name = "urb"
  app.Usage = "lookup urbandictionary.com from the command line. For instance: `urb 'fat beats'`"

  app.Flags = []cli.Flag {
      cli.BoolFlag{
        Name: "examples, e",
        Usage: "show definition examples",
      },
      cli.IntFlag{
        Name: "num, n",
        Value: 1,
        Usage: "number of results to show",
      },
    }

  app.Action = func(c *cli.Context) error {
    str := c.Args().Get(0)
    if str == "" {
      fmt.Println("Error: You must pass a search string. Run `urb help` for more info")
      return nil
    }

    resp, err := http.Get(urbandictionaryApi + url.QueryEscape(str))
    defer resp.Body.Close()
    if err != nil {
      fmt.Println("Error: Could not reach the Urban Dictionary API")
    	return nil
    }

    var res udResponse
    err = json.NewDecoder(resp.Body).Decode(&res)
    if err != nil {
      fmt.Println("Error: Could not decode the Urban Dictionary API response")
      return nil
    }

    printDefinitions(&res, c.Int("num"), c.Bool("examples"))

    return nil
  }

  app.Run(os.Args)
}

func printDefinitions(res *udResponse, num int, printExamples bool) {
  if num > len(res.List) { num = len(res.List) }
  for i := 0; i < num; i++ {
    def := res.List[i]
    if num == 1 {
      fmt.Println(def.Definition)
    } else {
      fmt.Printf("%d: " + def.Definition + "\n", i + 1)
    }
    if printExamples {
      fmt.Println("Example: '" + def.Example + "'\n")
    }
  }
}
