package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// V3Client Github API v3 client
type V3Client struct {
	Token      string // AccessToken
	Owner      string // Owner
	Repository string // Repository
}

// Authenticate 認証
func authenticate(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

// CreateIssue Issue作成
func (t V3Client) CreateIssue(title, body string) (*github.Issue, error) {
	request := github.IssueRequest{
		Title: &title,
		Body:  &body,
	}
	issue, resp, err := authenticate(t.Token).Issues.Create(context.Background(), t.Owner, t.Repository, &request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("create failed: %s", resp.Status)
	}
	return issue, nil
}

// EditIssue Issue更新
func (t V3Client) EditIssue(number int, title, body string) (*github.Issue, error) {
	request := github.IssueRequest{
		Title: &title,
		Body:  &body,
	}
	issue, resp, err := authenticate(t.Token).Issues.Edit(context.Background(), t.Owner, t.Repository, number, &request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("edit issue failed: %s", resp.Status)
	}
	return issue, nil
}

// CloseIssue Issueクローズ
func (t V3Client) CloseIssue(number int) (*github.Issue, error) {
	state := "close"
	request := github.IssueRequest{
		State: &state,
	}
	issue, resp, err := authenticate(t.Token).Issues.Edit(context.Background(), t.Owner, t.Repository, number, &request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("close issue failed: %s", resp.Status)
	}
	return issue, nil
}
