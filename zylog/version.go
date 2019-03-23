package zylog

import "fmt"

var (
	// Version is populated at compile time by govvv from ./VERSION
	Version string

	// GitCommit is populated at compile time by govvv.
	GitCommit string

	// GitState is populated at compile time by govvv.
	GitState string

	// GitBranch is current branch name the code is built off
	GitBranch string

	// BuildDate is RFC3339 formatted UTC date
	BuildDate string
)

func VersionString() string {
	if Version == "" {
		return "N/A"
	}
	return Version
}

func BuildString() string {
	if GitCommit == "" {
		return "N/A"
	}
	return fmt.Sprintf("%s@%s, %s", GitBranch, GitCommit, BuildDate)
}
