package cmd

import (
	"fmt"

	cli "github.com/spf13/cobra"

	"github.com/Ilyes512/boilr/pkg/boilr"
	"github.com/Ilyes512/boilr/pkg/util/validate"
)

// Version contains the cli-command for printing the current version of the tool.
var Version = &cli.Command{
	Use:   "version",
	Short: "Show the boilr version information",
	Run: func(cmd *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})

		shouldntPrettify := GetBoolFlag(cmd, "dont-prettify")
		if shouldntPrettify {
			fmt.Println(boilr.Version)
		} else {
			fmt.Println("Version: ", boilr.Version)
			fmt.Println("Build Date (UTC): ", boilr.BuildDate)
			fmt.Println("Commit:", boilr.Commit)
		}
	},
}
