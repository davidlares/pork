package pork

import (
  "github.com/davidlares/nap"
  "github.com/spf13/cobra"
  "encoding/json"
  "io/ioutil"
  "fmt"
  "net/http"
  "strings"
  "log"
)

type SearchResponse struct {
  Results []*SearchResult `json:"items"` // return another slice
}

type SearchResult struct {
  FullName string `json:"full_name"`
}

var SearchCmd = &cobra.Command{
  Use: "search", // subcommand
  Short: "Search for GH repos by keyword",
  Run: func(cmd *cobra.Command, args []string){
    if err := SearchByKeyword(args); err != nil {
      log.Fatalln("Search failed: ", err)
    }
  },
}

func SearchByKeyword(keywords []string) error {
  return GitHubAPI().Call("search", map[string]string {
    "query": strings.Join(keywords, "+"), // forming query
  }, nil)
}

// success response
func SearchSuccess(resp *http.Response) error {
  defer resp.Body.Close()
  content, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  respContent := SearchResponse{}
  json.Unmarshal(content, &respContent)
  for _, item := range respContent.Results {
    fmt.Println(item.FullName)
  }
  return nil
}

// check for code status
func SearchDefaultRouter(resp *http.Response) error {
  return fmt.Errorf("status code %d", resp.StatusCode)
}

//  responses
func GetSearchResource() *nap.RestResource {
  searchRouter := nap.NewRouter()
  searchRouter.DefaultRouter = SearchDefaultRouter // this is for non 200 statuses
  searchRouter.RegisterFunc(200, SearchSuccess)
  search := nap.NewResource("/search/repositories?q={{.query}}", "GET", searchRouter)
  return search
}
