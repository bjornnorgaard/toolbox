package repoedit

import (
	"fmt"

	"github.com/bjornnorgaard/toolbox/tools/github/types"
	"github.com/cli/go-gh"
)

const (
	off = "--flag=false"

	enableAutoMerge     = "--enable-auto-merge"
	deleteBranchOnMerge = "--delete-branch-on-merge"
	enableSquashMerge   = "--enable-squash-merge"
	showUpdateBranch    = "--allow-update-branch"

	disableAutoMerge   = "--enable-auto-merge" + " " + off
	keepBranchOnMerge  = "--delete-branch-on-merge" + " " + off
	disableSquashMerge = "--enable-squash-merge" + " " + off
	NoUpdateBranch     = "--allow-update-branch" + " " + off
)

type optsApply func(o *optsType)

type optsType struct {
	flags []string
}

var optsDefault = optsType{
	flags: []string{},
}

func WithEnableAutoMerge() optsApply {
	return func(o *optsType) {
		o.flags = append(o.flags, enableAutoMerge)
	}
}

func WithDeleteBranchOnMerge() optsApply {
	return func(o *optsType) {
		o.flags = append(o.flags, deleteBranchOnMerge)
	}
}

func WithEnableSquashMerge() optsApply {
	return func(o *optsType) {
		o.flags = append(o.flags, enableSquashMerge)
	}
}

func WithShowUpdateBranch() optsApply {
	return func(o *optsType) {
		o.flags = append(o.flags, showUpdateBranch)
	}
}

func WithDisableAutoMerge() optsApply {
	return func(o *optsType) {
		o.flags = append(o.flags, disableAutoMerge)
	}
}

func WithKeepBranchOnMerge() optsApply {
	return func(o *optsType) {
		o.flags = append(o.flags, keepBranchOnMerge)
	}
}

func WithDisableSquashMerge() optsApply {
	return func(o *optsType) {
		o.flags = append(o.flags, disableSquashMerge)
	}
}

func WithHiddenUpdateBranch() optsApply {
	return func(o *optsType) {
		o.flags = append(o.flags, NoUpdateBranch)
	}
}

func Update(repo types.Repo, appliers ...optsApply) error {
	opts := &optsDefault
	for _, applier := range appliers {
		applier(opts)
	}

	if len(opts.flags) == 0 {
		return fmt.Errorf("no options specified")
	}

	command := []string{"repo", "edit", repo.FullName}
	command = append(command, opts.flags...)

	if _, _, err := gh.Exec(command...); err != nil {
		return fmt.Errorf("failed to enable auto merge for repo '%s': %w", repo.FullName, err)
	}

	return nil
}
