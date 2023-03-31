package main

import (
	"fmt"
	"os"

  "github.com/nhudson/polybar-gh-notification/internal/github"
	"github.com/spf13/cobra"
)

var (
  githubToken string
  interval int
)

func main() {
  cmd := cobra.Command{
    Use: "poly-gh-notifier [cmd]",
    Short: "Poll Github API for notifications",
    RunE: func(cmd *cobra.Command, args []string) error {
      if githubToken == "" {
        return fmt.Errorf("github token is required")
      }
      return github.Run(githubToken, interval)
    },
  }

  cmd.Flags().StringVarP(&githubToken, "git-token", "t", os.Getenv("GITHUB_TOKEN"), "Github token")
  cmd.Flags().IntVarP(&interval, "interval", "i", 10, "Polling interval")

  if err := cmd.Execute(); err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }
}
