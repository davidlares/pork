package pork

import (
  "testing"
  "github.com/davidlares/nap"
)
func TestGetRepositoryReadme(t *testing.T) {
  token := "564864432131655181615165168"
  GitHubAPI().SetAuth(nap.NewAuthToken(token))
  if err := GetRepositoryReadme("davidlares/testrepo"); err != nil {
    t.Fail()
  }
}
