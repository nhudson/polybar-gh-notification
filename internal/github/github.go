package github

import (
  "context"
  "fmt"
  "os"
  "time"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubClient struct {
  client *github.Client
}

func Run(token string, interval int) error {
  // Setup Github client
  gh := NewGithubClient(token, context.Background())
  for {
    notifications, err := gh.GetNotifications(context.Background())
    if err != nil {
      fmt.Println("Error: ", err)
      os.Exit(1)
    }

    // Get a count of how many notifications are unread
    unread := 0
    for _, notification := range notifications {
      if *notification.Unread {
        unread++
      }
    }

    // Print the count
    fmt.Printf("  : %d\n", unread)

    time.Sleep(time.Second * time.Duration(interval))
  }
}

func NewGithubClient(token string, ctx context.Context) *GithubClient {
  return &GithubClient{
    client: github.NewClient(
      oauth2.NewClient(
        ctx,
        oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
      ),
    ),
  }
}

func (gc *GithubClient) GetNotifications(ctx context.Context) ([]*github.Notification, error) {

  // Set ListNotificationsOptions
  opt := &github.NotificationListOptions{
    All: true,
  }

  notifications, _, err := gc.client.Activity.ListNotifications(ctx, opt)
  if err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }

  return notifications, nil
}
