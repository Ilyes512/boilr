package cmd

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"

	"github.com/Ilyes512/boilr/pkg/boilr"
	"github.com/Ilyes512/boilr/pkg/host"
	"github.com/Ilyes512/boilr/pkg/util/exit"
	"github.com/Ilyes512/boilr/pkg/util/git"
	"github.com/Ilyes512/boilr/pkg/util/osutil"
	"github.com/Ilyes512/boilr/pkg/util/tlog"
	"github.com/Ilyes512/boilr/pkg/util/validate"
)

// Download contains the cli-command for downloading templates from github.
var Download = &cli.Command{
	Use:   "download <template-repo> <template-tag>",
	Short: "Download a project template from a github repository to template registry",
	// FIXME Half-Updates leave messy templates
	Run: func(cmd *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-repo", validate.UnixPath},
			{"template-tag", validate.AlphanumericExt},
		})

		MustValidateTemplateDir()

		tlog.SetLogLevel(GetStringFlag(cmd, "log-level"))

		templateURL, templateName := args[0], args[1]

		targetDir, err := boilr.TemplatePath(templateName)
		if err != nil {
			exit.Error(fmt.Errorf("download: %s", err))
		}

		switch exists, err := osutil.DirExists(targetDir); {
		case err != nil:
			exit.Error(fmt.Errorf("download: %s", err))
		case exists:
			if shouldOverwrite := GetBoolFlag(cmd, "force"); !shouldOverwrite {
				exit.OK("Template %v already exists use -f to overwrite the template", templateName)
			}

			// TODO(tmrts): extract `template delete` helper and use that one
			if err := os.RemoveAll(targetDir); err != nil {
				exit.Error(fmt.Errorf("download: %s", err))
			}
		}

		// TODO(tmrts): allow fetching other branches than 'master'
		if err := git.Clone(targetDir, git.CloneOptions{
			URL: host.URL(templateURL),
		}); err != nil {
			exit.Error(fmt.Errorf("download: %s", err))
		}

		// TODO(tmrts): use git-notes as metadata storage or boltdb
		if err := serializeMetadata(templateName, templateURL, targetDir); err != nil {
			exit.Error(fmt.Errorf("download: %s", err))
		}

		exit.OK("Successfully downloaded the template %#v", templateName)
	},
}
