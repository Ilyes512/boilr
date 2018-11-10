package cmd

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"

	"github.com/Ilyes512/boilr/pkg/boilr"
	"github.com/Ilyes512/boilr/pkg/util/exit"
	"github.com/Ilyes512/boilr/pkg/util/validate"
)

// Rename contains the cli-command for renaming templates in the template registry.
var Rename = &cli.Command{
	Hidden: true,
	Use:    "rename <old-template-tag> <new-template-tag>",
	Short:  "Rename a project template",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{Name: "old-template-tag", Validate: validate.UnixPath},
			{Name: "new-template-tag", Validate: validate.UnixPath},
		})

		MustValidateTemplateDir()

		tmplName, newTmplName := args[0], args[1]

		if ok, err := TemplateInRegistry(tmplName); err != nil {
			exit.Fatal(fmt.Errorf("rename: %s", err))
		} else if !ok {
			exit.Fatal(fmt.Errorf("Template %q couldn't be found in the template registry", tmplName))
		}

		tmplPath, err := boilr.TemplatePath(tmplName)
		if err != nil {
			exit.Fatal(fmt.Errorf("rename: %s", err))
		}

		newTmplPath, err := boilr.TemplatePath(newTmplName)
		if err != nil {
			exit.Fatal(fmt.Errorf("rename: %s", err))
		}

		if err := os.Rename(tmplPath, newTmplPath); err != nil {
			exit.Error(fmt.Errorf("rename: %s", err))
		}

		exit.OK("Successfully renamed the template %q to %q", tmplName, newTmplName)
	},
}
