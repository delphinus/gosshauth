package gosshauth

import (
	"fmt"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

// RepoName is a name of this product.
const RepoName = "delphinus/gosshauth"

// Version is needed for selfupdate.
const Version = "0.0.1"

// GitCommit is needed for selfupdate.
var GitCommit = ""

func doSelfupdate() error {
	v := semver.MustParse(Version)
	latest, err := selfupdate.UpdateSelf(v, RepoName)
	if err != nil {
		return fmt.Errorf("binary update failed: %v", err)
	}
	if latest.Version.Equals(v) {
		fmt.Println("current binary is the latest version", Version)
	} else {
		fmt.Println("successfully updated to version", latest.Version)
		fmt.Println("release note:\n", latest.ReleaseNotes)
	}
	return nil
}

func showVersion() error {
	fmt.Printf("%s %s (%s)\n", os.Args[0], Version, GitCommit)
	return nil
}
