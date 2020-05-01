package pork

// api configuration

import (
  "github.com/davidlares/nap"
  "github.com/spf13/viper"
)

var api *nap.API

// singleton -> called and returns and object -> the same object is returned
func GitHubAPI() *nap.API {
  if api == nil {
    api = nap.NewAPI("https://api.github.com") // base URL
    token := viper.GetString("token") // from config file
    api.SetAuth(nap.NewAuthToken(token))
    api.AddResource("fork", GetForkResource())
    api.AddResource("search", GetSearchResource())
    api.AddResource("docs", GetReadmeResource())
    api.AddResource("pullrequest", GetPullRequestResource())

  }
  return api
}
