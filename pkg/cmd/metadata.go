package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Ilyes512/boilr/pkg/boilr"
	"github.com/Ilyes512/boilr/pkg/template"
)

func serializeMetadata(tag string, repo string, targetDir string) error {
	fname := filepath.Join(targetDir, boilr.TemplateMetadataName)

	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)

	t := template.Metadata{Tag: tag, Repository: repo, Created: template.NewTime()}
	if err := enc.Encode(&t); err != nil {
		return err
	}

	return nil
}
