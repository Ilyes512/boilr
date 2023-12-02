package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	cli "github.com/spf13/cobra"

	"github.com/Ilyes512/boilr/pkg/boilr"
	"github.com/Ilyes512/boilr/pkg/template"
	"github.com/Ilyes512/boilr/pkg/util/exit"
	"github.com/Ilyes512/boilr/pkg/util/osutil"
	"github.com/Ilyes512/boilr/pkg/util/tlog"
	"github.com/Ilyes512/boilr/pkg/util/validate"
)

// TemplateInRegistry checks whether the given name exists in the template registry.
func TemplateInRegistry(name string) (bool, error) {
	names, err := ListTemplates()
	if err != nil {
		return false, err
	}

	_, ok := names[name]
	return ok, nil
}

// Use contains the cli-command for using templates located in the local template registry.
// TODO add --use-cache flag to execute a template from previous answers to prompts
var Use = &cli.Command{
	Use:   "use <template-tag> <target-dir>",
	Short: "Execute a project template in the given directory",
	Run: func(cmd *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{Name: "template-tag", Validate: validate.UnixPath},
			{Name: "target-dir", Validate: validate.UnixPath},
		})

		MustValidateTemplateDir()

		tlog.SetLogLevel(GetStringFlag(cmd, "log-level"))

		tmplName := args[0]
		targetDir, err := filepath.Abs(args[1])
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		templateFound, err := TemplateInRegistry(tmplName)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		if !templateFound {
			exit.Fatal(fmt.Errorf("template %q couldn't be found in the template registry", tmplName))
		}

		tmplPath, err := boilr.TemplatePath(tmplName)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		tmpl, err := template.Get(tmplPath)
		if err != nil {
			exit.Fatal(fmt.Errorf("use: %s", err))
		}

		if shouldUseDefaults := GetBoolFlag(cmd, "use-defaults"); shouldUseDefaults {
			tmpl.UseDefaultValues()
		}

		executeTemplate := func() error {
			parentDir := filepath.Dir(targetDir)

			exists, err := osutil.DirExists(parentDir)
			if err != nil {
				return err
			}

			if !exists {
				return fmt.Errorf("use: parent directory %q doesn't exist", parentDir)
			}

			tmpDir, err := os.MkdirTemp("", "boilr-use-template")
			if err != nil {
				return err
			}
			defer os.RemoveAll(tmpDir)

			if err := tmpl.Execute(tmpDir); err != nil {
				return err
			}

			// Complete the template execution transaction by copying the temporary dir to the target directory.
			return osutil.CopyRecursively(tmpDir, targetDir)
		}

		if err := executeTemplate(); err != nil {
			exit.Fatal(fmt.Errorf("use: %v", err))
		}

		exit.OK("Successfully executed the project template %v in %v", tmplName, targetDir)
	},
}
