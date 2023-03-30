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
    _, err := gh.GetNotifications(context.Background())
    if err != nil {
      fmt.Println("Error: ", err)
      os.Exit(1)
    }

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

func (c *GithubClient) GetNotifications(ctx context.Context) ([]*github.Notification, error) {

  // Set ListNotificationsOptions
  opt := &github.NotificationListOptions{
    ListOptions: github.ListOptions{PerPage: 10},
  }
  notifications, _, err := c.client.Activity.ListNotifications(ctx, opt)
  if err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }

  // Print GetNotifications 
  for _, notification := range notifications {
    fmt.Println("Notification: ", notification)
  }

  return notifications, nil
}
