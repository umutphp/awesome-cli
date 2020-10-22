package selfupdate

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const gitHubRepo = "umutphp/awesome-cli"

func Update(currentVersion string) error {
	v, err := semver.Parse(currentVersion)
	if err != nil {
		return err
	}

	latest, err := selfupdate.UpdateSelf(v, gitHubRepo)
	if err != nil {
		return err
	}

	if latest.Version.LTE(v) {
		fmt.Printf(
			"awesome-cli is up-to-date: v%s\nGo forth and be awesome!\n", v)

		return nil
	}

	fmt.Printf("awesome-cli updated to v%s\n", latest.Version)

	return nil
}
