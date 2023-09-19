package repos

import (
	"encoding/json"
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/types"
	"github.com/bjornnorgaard/toolbox/tools/github/user"
	"github.com/bjornnorgaard/toolbox/utils/jqexp"
	"github.com/cli/go-gh"
	"strings"
)

type optsType struct {
	owner string
	limit uint
}

var (
	fields = []string{
		"createdAt",
		"description",
		"fullName",
		"id",
		"name",
		"owner",
		"url",
		"visibility",
	}

	jq = jqexp.New(`[.[] | {
		createdAt,
		description,
		fullName,
		id,
		name,
		owner: .owner.login,
		ownerId: .owner.id,
		url,
		visibility,
	}]`)

	optsDefault = optsType{
		limit: 1000,
	}
)

func GetRepos(applies ...optsApply) ([]types.Repo, error) {
	opts := &optsDefault
	for _, apply := range applies {
		apply(opts)
	}

	buffer, _, err := gh.Exec("search", "repos",
		fmt.Sprintf("--limit=%d", opts.limit),
		fmt.Sprintf("--owner=%s", user.Me()),
		fmt.Sprintf("--json=%s", strings.Join(fields, ",")),
		fmt.Sprintf("--jq=%s", jq),
	)

	if err != nil {
		return nil, err
	}

	var repos []types.Repo
	if err = json.Unmarshal(buffer.Bytes(), &repos); err != nil {
		return nil, fmt.Errorf("failed to unmarshal repositories '%+v': %w", buffer.String(), err)
	}

	return repos, nil
}

type optsApply func(o *optsType)

func withLimit(limit uint) optsApply {
	return func(o *optsType) {
		o.limit = limit
	}
}
