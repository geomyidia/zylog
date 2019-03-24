package logger

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

// Returns a version string as set by the Makefile when the library was last
// compiled. If a version cannot be extracted, the string 'NA' is returned.
func VersionString() string {
	if Version == "" {
		return "N/A"
	}
	return Version
}

// Returns a build string that provides the git branch upon which the build was
// made, the git commit of that branch, as well as the build date. If a build
// string cannot be constructed, the string 'NA' is returned.
func BuildString() string {
	if GitCommit == "" {
		return "N/A"
	}
	return fmt.Sprintf("%s@%s, %s", GitBranch, GitCommit, BuildDate)
}
