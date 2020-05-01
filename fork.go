package pork

import (
  "github.com/davidlares/nap"
  "github.com/spf13/cobra"
  "encoding/json"
  "net/http"
  "io/ioutil"
  "log"
  "fmt"
  "strings"
)

type ForkResponse struct {
  CloneURL string `json:"clone_url"`
  FullName string `json"full_name"`
}

var ForkCmd = &cobra.Command {
  Use: "fork",
  Short: "fork a Github repository",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) <= 0 {
      log.Fatalln("you must supply a repository")
    }
    if err := ForkRepository(args[0]); err != nil {
      log.Fatalln("unable to fork repository: ", err)
    }
  },
}

// fork functionality
func ForkRepository(repository string) error {
  values := strings.Split(repository, "/")
  if len(values) != 2 {
    return fmt.Errorf("Repository format in invalid - owner/repo")
  }
  return GitHubAPI().Call("fork", map[string]string{
    "owner": values[0],
    "repo": values[1],
  }, nil)
}

// success
func ForkSuccess(resp *http.Response) error {
  defer resp.Body.Close()
  content, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }
  respContent := ForkResponse{}
  json.Unmarshal(content, &respContent)
  fmt.Printf("Forked to repository: %s \n", respContent.FullName)
  return nil
}

// resource for the NAP library
func GetForkResource() *nap.RestResource {
  forkRouter := nap.NewRouter()
  forkRouter.RegisterFunc(202, ForkSuccess)
  forkRouter.RegisterFunc(404, func(_ *http.Response) error {
    return fmt.Errorf("You must set an authentication token")
  })
  fork := nap.NewResource("/repos/{{.owner}}/{{.repo}}/forks","POST", forkRouter)
  return fork
}
