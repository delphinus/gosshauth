package gosshauth

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"gopkg.in/urfave/cli.v2"
)

// RepoName is a name of this product.
const RepoName = "delphinus/gosshauth"

// Selfupdate is a command action for `selfupdate`.
func Selfupdate(c *cli.Context) error {
	latest, found, err := selfupdate.DetectLatest(RepoName)
	if err != nil {
		return fmt.Errorf("error in DetectLatest: %v", err)
	}
	v := semver.MustParse(Version)
	if !found || latest.Version.LTE(v) {
		fmt.Fprintln(c.App.Writer, "current binary is the latest version", Version)
		return nil
	}
	fmt.Fprintln(c.App.Writer, "updating...", v, "=>", latest.Version)
	if err := selfupdate.UpdateTo(latest.AssetURL, Me()); err != nil {
		return fmt.Errorf("binary update failed: %v", err)
	}
	fmt.Fprintln(c.App.Writer, "successfully updated to version", latest.Version)
	fmt.Fprintln(c.App.Writer, "release note:\n", latest.ReleaseNotes)
	return nil
}
