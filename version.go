package tst

import (
	"fmt"
)

// Variables set by ldflags representing build version.
var (
	BuildDate = "<not set>"

	// ScmRev represents source control management revision, e.g.: v1.2,
	ScmRev = "<not set>"

	// ScmHash represents source control management hash, e.g.: 12ab23,
	ScmHash = "<not set>"

	// ScmWDState represents working directory state: clean, dirty.
	ScmWDState = "<not set>"

	// CCTag represents CI/CD tag.
	CCTag = "<not set>"
)

// VersionString returns version string with given command name.
func VersionString(cmd string) string {
	format := "%s %s, hash: %s, build date: %s"
	return fmt.Sprintf(format, cmd, ScmRev, ScmHash, BuildDate)
}
