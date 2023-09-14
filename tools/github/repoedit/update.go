package repoedit

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/types"
	"github.com/cli/go-gh"
)

func Update(repo types.Repo, appliers ...OptsApply) error {
	defaultOpts := DefaultOpts()
	opts := &defaultOpts

	for _, applier := range appliers {
		applier(opts)
	}

	flags := make([]string, 0, len(opts.Settings))
	for setting, enabled := range opts.Settings {
		flags = append(flags, fmt.Sprintf("%s=%t", setting, enabled))
	}

	command := []string{"repo", "edit", repo.FullName}
	command = append(command, flags...)

	if !opts.Debug {
		_, _, err := gh.Exec(command...)
		if err != nil {
			return err
		}
	}

	return nil
}

func DefaultOpts() OptsType {
	return OptsType{
		Debug: false,
		Settings: map[RepoSetting]bool{
			SettingDeleteBranchOnMerge: true,
			SettingEnableSquashMerge:   true,
			SettingEnableRebaseMerge:   true,
			SettingEnableAutoMerge:     true,
			SettingEnableIssues:        true,
			SettingAllowUpdateBranch:   true,
			SettingEnableDiscussions:   false,
			SettingEnableMergeCommit:   false,
			SettingEnableProjects:      false,
			SettingEnableWiki:          false,
		},
	}
}

type RepoSetting string

const (
	// SettingAllowUpdateBranch => Allow a pull request head branch that is behind its base branch to be updated
	SettingAllowUpdateBranch RepoSetting = "--allow-update-branch"

	// SettingDeleteBranchOnMerge => Delete head branch when pull requests are merged
	SettingDeleteBranchOnMerge RepoSetting = "--delete-branch-on-merge"

	// SettingEnableAutoMerge => Enable auto-merge functionality
	SettingEnableAutoMerge RepoSetting = "--enable-auto-merge"

	// SettingEnableDiscussions => Enable discussions in the repository
	SettingEnableDiscussions RepoSetting = "--enable-discussions"

	// SettingEnableIssues => Enable issues in the repository
	SettingEnableIssues RepoSetting = "--enable-issues"

	// SettingEnableMergeCommit => Enable merging pull requests via merge commit
	SettingEnableMergeCommit RepoSetting = "--enable-merge-commit"

	// SettingEnableProjects => Enable projects in the repository
	SettingEnableProjects RepoSetting = "--enable-projects"

	// SettingEnableRebaseMerge => Enable merging pull requests via rebase
	SettingEnableRebaseMerge RepoSetting = "--enable-rebase-merge"

	// SettingEnableSquashMerge => Enable merging pull requests via squashed commit
	SettingEnableSquashMerge RepoSetting = "--enable-squash-merge"

	// SettingEnableWiki => Enable wiki in the repository
	SettingEnableWiki RepoSetting = "--enable-wiki"
)

type OptsType struct {
	Debug    bool
	Settings map[RepoSetting]bool
}

type OptsApply func(o *OptsType)

func WithDebug(enabled bool) OptsApply {
	return func(o *OptsType) {
		o.Debug = enabled
	}
}

func With(setting RepoSetting, enabled bool) OptsApply {
	return func(o *OptsType) {
		o.Settings[setting] = enabled
	}
}
