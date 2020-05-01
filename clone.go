package pork

import (
  "github.com/spf13/cobra"
  "fmt"
  "github.com/spf13/viper"
  "log"
)

var CloneCmd = &cobra.Command {
  Use: "clone",
  Short: "Clone repository from GH",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) <= 0 {
      log.Fatalln("You must supply the repository")
    }
    if err := CloneRepository(args[0], ref, create); err != nil {
      log.Fatalln("error when closing repository: ", err)
    }
  },
}

var ref string
var create bool

func CloneRepository(repository, ref string, shouldCreate bool) error {
  // creating new repo
  repo, err := NewGHRepo(repository)
  if err != nil {
    return err
  }
  if err := repo.Clone(viper.GetString("location")); err != nil {
    return err
  }
  if err := repo.Checkout(ref, shouldCreate); err != nil {
    return err
  }
  fmt.Printf("Cloned repository to: %s \n", repo.RepoDir)
  return nil
}

func init() {
  // config options flags
  CloneCmd.PersistentFlags().StringVar(&ref, "ref", "master", "specific reference to check out") // master by default
  CloneCmd.PersistentFlags().BoolVar(&create, "create", false, "create the reference if it does not exists")
}
