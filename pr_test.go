package pork

import (
  "testing"
  "github.com/davidlares/nap"
)

func TestPullRequest(t *testing.T) {
  token := "564864432131655181615165168"
  GitHubAPI().SetAuth(nap.NewAuthToken(token))
  destRepo = "davidlares/testrepo:master"
  sourceRepo = "davidlares/testrepo:master"
  pullRequestTitle = "test pull request"
  pullRequestMessage = "here it is"
  if err := CreatePullRequest(); err != nil {
    t.Fail()
  }
}
