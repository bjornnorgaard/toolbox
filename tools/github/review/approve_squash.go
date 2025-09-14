package review

import (
	"fmt"

	"github.com/bjornnorgaard/toolbox/tools/github/types"
	"github.com/cli/go-gh"
)

const (
	dependabotSquash   = "@dependabot squash and merge"
	dependabotRecreate = "@dependabot recreate"
)

func ApproveSquash(pr types.PR, opts ...ApplyOpts) error {
	defaultOpt := OptsType{
		messageBody: dependabotSquash,
	}
	for _, apply := range opts {
		apply(&defaultOpt)
	}

	_, _, err := gh.Exec("pr", "review",
		fmt.Sprintf("%d", pr.Number),
		fmt.Sprintf("%s", "--approve"),
		fmt.Sprintf("--body=%s", defaultOpt.messageBody),
		fmt.Sprintf("--repo=%s", pr.RepositoryWithOwner),
	)

	if err != nil {
		return err
	}

	return nil
}

type OptsType struct {
	messageBody string
}

type ApplyOpts func(o *OptsType)

func WithSquash() ApplyOpts {
	return func(o *OptsType) {
		o.messageBody = dependabotSquash
	}
}

func WithRecreate() ApplyOpts {
	return func(o *OptsType) {
		o.messageBody = dependabotRecreate
	}
}
