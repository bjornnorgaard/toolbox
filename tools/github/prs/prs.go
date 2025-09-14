package prs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bjornnorgaard/toolbox/tools/github/types"
	"github.com/bjornnorgaard/toolbox/tools/github/user"
	"github.com/bjornnorgaard/toolbox/utils/jqexp"
	"github.com/cli/go-gh"
)

type OptsApply func(o *optsType)

type optsType struct {
	app    string
	checks string
	limit  uint
	owner  string
	review string
	state  string
}

func getDefaultOpts() *optsType {
	return &optsType{
		app:    "dependabot",
		checks: "success",
		limit:  100,
		owner:  user.Me(),
		review: "none",
		state:  "open",
	}
}

func WithStateOpen() OptsApply {
	return func(o *optsType) {
		o.state = "open"
	}
}

func WithStateClosed() OptsApply {
	return func(o *optsType) {
		o.state = "closed"
	}
}

func WithAuthorBot() OptsApply {
	return func(o *optsType) {
		o.app = "dependabot"
	}
}

func WithReviewApproved() OptsApply {
	return func(o *optsType) {
		o.review = "approved"
	}
}

func WithReviewRequired() OptsApply {
	return func(o *optsType) {
		o.review = "required"
	}
}

func WithReviewChangesRequested() OptsApply {
	return func(o *optsType) {
		o.review = "changes_requested"
	}
}

func WithReviewNone() OptsApply {
	return func(o *optsType) {
		o.review = "none"
	}
}

func WithLimit(limit uint) OptsApply {
	return func(o *optsType) {
		o.limit = limit
	}
}

func WithChecksSucceeded() OptsApply {
	return func(o *optsType) {
		o.checks = "success"
	}
}

func WithChecksFailed() OptsApply {
	return func(o *optsType) {
		o.checks = "failure"
	}
}

func WithChecksPending() OptsApply {
	return func(o *optsType) {
		o.checks = "pending"
	}
}

var (
	fields = []string{
		"author",
		"createdAt",
		"id",
		"number",
		"repository",
		"state",
		"title",
		"updatedAt",
	}
	jq = jqexp.New(`[.[] | {
		author: .author.login,
		createdAt,
		id,
		number,
		repository: .repository.name,
		repositoryWithOwner: .repository.nameWithOwner,
		state,
		title,
		updatedAt,
	}]`)
)

func Get(applies ...OptsApply) ([]types.PR, error) {
	opts := getDefaultOpts()
	for _, apply := range applies {
		apply(opts)
	}

	buf, _, err := gh.Exec("search", "prs",
		fmt.Sprintf("--app=%s", opts.app),
		fmt.Sprintf("--checks=%s", opts.checks),
		fmt.Sprintf("--limit=%d", opts.limit),
		fmt.Sprintf("--owner=%s", opts.owner),
		fmt.Sprintf("--review=%s", opts.review),
		fmt.Sprintf("--state=%s", opts.state),
		fmt.Sprintf("--json=%s", strings.Join(fields, ",")),
		fmt.Sprintf("--jq=%s", jq),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch pull requests, config %+v: %w", opts, err)
	}

	var prs []types.PR
	if err = json.Unmarshal(buf.Bytes(), &prs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal pull requests '%+v': %w", buf.String(), err)
	}

	return prs, nil
}
