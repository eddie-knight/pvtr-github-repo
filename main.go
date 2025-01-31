package main

import (
	"fmt"

	"os"

	"github.com/eddie-knight/pvtr-osps-baseline/armory"

	"github.com/privateerproj/privateer-sdk/command"
	"github.com/privateerproj/privateer-sdk/config"
)

var (
	// Version is to be replaced at build time by the associated tag
	Version = "0.0.0"
	// VersionPostfix is a marker for the version such as "dev", "beta", "rc", etc.
	VersionPostfix = "dev"
	// GitCommitHash is the commit at build time
	GitCommitHash = ""
	// BuiltAt is the actual build datetime
	BuiltAt = ""

	PluginName   = "github-repo"
	RequiredVars = []string{
		"owner",
		"repo",
	}

	runCmd = command.NewPluginCommands(
		PluginName,
		Version,
		VersionPostfix,
		GitCommitHash,
		&armory.Armory,
		initializer,
		RequiredVars,
	)
)

// initializer is a custom function to set up the armory for our usecase
func initializer(c *config.Config) (err error) {
	armory.SetupArmory(c)
	return
}

func main() {
	if VersionPostfix != "" {
		Version = fmt.Sprintf("%s-%s", Version, VersionPostfix)
	}

	err := runCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
