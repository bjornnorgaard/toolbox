package prs

import (
	"encoding/json"
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/curuser"
	"github.com/bjornnorgaard/toolbox/tools/github/types"
	"github.com/bjornnorgaard/toolbox/utils/jqexp"
	"github.com/cli/go-gh"
	"strings"
)

type optsApply func(o *optsType)

type optsType struct {
	author string
	state  string
	checks string
	review string
	limit  uint
}

var optsDefault = optsType{
	author: "app/dependabot",
	state:  "open",
	checks: "success",
	review: "none",
	limit:  1000,
}

func WithStateOpen() optsApply {
	return func(o *optsType) {
		o.state = "open"
	}
}

func WithStateClosed() optsApply {
	return func(o *optsType) {
		o.state = "closed"
	}
}

func WithAuthorBot() optsApply {
	return func(o *optsType) {
		WithAuthor("app/dependabot")
	}
}

func WithAuthor(author string) optsApply {
	return func(o *optsType) {
		o.author = author
	}
}

func WithReviewApproved() optsApply {
	return func(o *optsType) {
		o.review = "approved"
	}
}

func WithReviewRequired() optsApply {
	return func(o *optsType) {
		o.review = "required"
	}
}

func WithReviewChangesRequested() optsApply {
	return func(o *optsType) {
		o.review = "changes_requested"
	}
}

func WithReviewNotApproved() optsApply {
	return func(o *optsType) {
		o.review = "none"
	}
}

func WithLimit(limit uint) optsApply {
	return func(o *optsType) {
		o.limit = limit
	}
}

func WithChecksSucceeded() optsApply {
	return func(o *optsType) {
		o.checks = "success"
	}
}

func WithChecksFailed() optsApply {
	return func(o *optsType) {
		o.checks = "failure"
	}
}

func WithChecksPending() optsApply {
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

func Get(applies ...optsApply) ([]types.PR, error) {
	opts := &optsDefault
	for _, apply := range applies {
		apply(opts)
	}

	buf, _, err := gh.Exec("search", "prs",
		fmt.Sprintf("--author=%s", opts.author),
		fmt.Sprintf("--state=%s", opts.state),
		fmt.Sprintf("--review=%s", opts.review),
		fmt.Sprintf("--limit=%d", opts.limit),
		fmt.Sprintf("--owner=%s", curuser.Me()),
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
