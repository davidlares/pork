package pork

import (
  "testing"
  "github.com/davidlares/nap"
)

func TestSearchByKeyword(t *testing.T) {
  token := "564864432131655181615165168"
  GitHubAPI().SetAuth(nap.NewAuthToken(token))
  if err := SearchByKeyword([]string {"topic:go"}); err != nil {
    t.Fail()
  }
}
