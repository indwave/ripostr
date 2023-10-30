package ripostr

import "fmt"

const (
	// URL is the git URL for the repository.
	URL = "github.com/wave/ripostr"
	// GitRef is the gitref, as in refs/heads/branchname.
	GitRef = ""
	// ParentGitCommit is the commit hash of the parent HEAD.
	ParentGitCommit = ""
	// BuildTime stores the time when the current binary was built.
	BuildTime = ""
	// SemVer lists the (latest) git tag on the release.
	SemVer = "v0.0.1"
	// Major is the major number from the tag.
	Major = 0
	// Minor is the minor number from the tag.
	Minor = 0
	// Patch is the patch version number from the tag.
	Patch = 1
)

var CI = "false"

// Version returns a pretty printed version information string.
func Version() string {
	return fmt.Sprint(
		"\nRepository Information\n",
		"\tGit repository: "+URL+"\n",
		"\tBranch: "+GitRef+"\n",
		"\tParentGitCommit: "+ParentGitCommit+"\n",
		"\tBuilt: "+BuildTime+"\n",
		"\tSemVer: "+SemVer+"\n",
	)
}
