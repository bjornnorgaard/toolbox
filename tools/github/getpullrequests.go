package github

import (
	"encoding/json"
	"fmt"
	"github.com/cli/go-gh"
	"strings"
)

// getPullRequestsV1 fetches all pull requests using passed query filter.
func getPullRequestsV1(repo string, filter string) ([]PullRequestsV1, error) {
	query := fmt.Sprintf("is:open is:pr %s", filter)

	buffer, _, err := gh.Exec("pr", "list", "--repo", repo, "--search", query, "--json", "number,title")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pull-requests for repo %s - %v", repo, err)
	}

	var pullRequests []PullRequestsV1
	err = json.Unmarshal(buffer.Bytes(), &pullRequests)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal pull-requests for %s - %v", repo, err)
	}
	return pullRequests, nil
}

const (
	stateOpen   = "open"
	stateClosed = "closed"

	authorBot = "app/dependabot"

	reviewNone     = "none"
	reviewRequired = "required"
	reviewApproved = "approved"

	checkSuccess = "success"
	checkFailure = "failure"
	checkPending = "pending"
)

type SearchOptsFunc func(o *getPullOption)

func isOpen() SearchOptsFunc {
	return func(o *getPullOption) {
		o.state = stateOpen
	}
}

func isClosed() SearchOptsFunc {
	return func(o *getPullOption) {
		o.state = stateClosed
	}
}

func fromBot() SearchOptsFunc {
	return func(o *getPullOption) {
		fromAuthor(authorBot)
	}
}

func fromAuthor(author string) SearchOptsFunc {
	return func(o *getPullOption) {
		o.author = author
	}
}

func isApproved() SearchOptsFunc {
	return func(o *getPullOption) {
		o.review = reviewApproved
	}
}

func isNotApproved() SearchOptsFunc {
	return func(o *getPullOption) {
		o.review = reviewNone
	}
}

func withLimit(limit uint) SearchOptsFunc {
	return func(o *getPullOption) {
		o.limit = limit
	}
}

func isSucceeding() SearchOptsFunc {
	return func(o *getPullOption) {
		o.checks = checkSuccess
	}
}

type getPullOption struct {
	author string
	owner  string
	state  string
	checks string
	review string
	limit  uint
	fields []string
	jq     string
}

var defaultGetPullOption = getPullOption{
	author: authorBot,
	owner:  "@me",
	state:  stateOpen,
	checks: checkSuccess,
	review: reviewNone,
	limit:  100,
	fields: []string{
		"author",
		"createdAt",
		"id",
		"number",
		"repository",
		"state",
		"title",
		"updatedAt",
	},
	jq: strings.ReplaceAll(strings.ReplaceAll(`[.[] | {
		author: .author.login,
		createdAt,
		id,
		number,
		repository: .repository.name,
		repositoryWithOwner: .repository.nameWithOwner,
		state,
		title,
		updatedAt
	}]`, "\n", ""), "\t", ""),
}

func getprs(opts ...SearchOptsFunc) ([]PullRequest, error) {
	o := &defaultGetPullOption
	for _, optFunc := range opts {
		optFunc(o)
	}

	buf, _, err := gh.Exec("search", "prs",
		fmt.Sprintf("--owner=%s", o.owner),
		fmt.Sprintf("--author=%s", o.author),
		fmt.Sprintf("--state=%s", o.state),
		fmt.Sprintf("--review=%s", o.review),
		fmt.Sprintf("--limit=%d", o.limit),
		fmt.Sprintf("--json=%s", strings.Join(o.fields, ",")),
		fmt.Sprintf("--jq=%s", o.jq),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch prs, config %+v: %w", o, err)
	}

	var prs []PullRequest
	if err = json.Unmarshal(buf.Bytes(), &prs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal prs '%+v': %w", buf.String(), err)
	}

	return prs, nil
}
