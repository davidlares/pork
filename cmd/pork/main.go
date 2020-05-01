package main

import (
  "github.com/spf13/cobra"
  "github.com/davidlares/pork"
  "github.com/spf13/viper"
  "os"
  "fmt"
)

var rootCmd *cobra.Command

func main() {
  rootCmd.Execute()
}

func init() {
  rootCmd = &cobra.Command{
    Use: "pork",
    Short: "Project Forking for GH", // short description
  }
  rootCmd.AddCommand(pork.SearchCmd)
  rootCmd.AddCommand(pork.DocsCmd)
  rootCmd.AddCommand(pork.CloneCmd)
  rootCmd.AddCommand(pork.ForkCmd)
  rootCmd.AddCommand(pork.PullRequestCmd)

  viper.SetDefault("location", os.Getenv("HOME")) // repository stored location
  viper.SetConfigName("pork")
  viper.AddConfigPath(".")
  // check location
  if err := viper.ReadInConfig(); err != nil {
    fmt.Println("No config file found")
  }
  // setting default
  viper.SetDefault("location", os.Getenv("HOME")) // $HOME is not found by default
}
