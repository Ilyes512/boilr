package cmd_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Ilyes512/boilr/pkg/cmd"
)

func TestUseExecutesProjectTemplate(t *testing.T) {
	t.SkipNow()
	tmpDirPath, err := os.MkdirTemp("", "template-test")
	if err != nil {
		t.Fatalf("os.MkdirTemp() got error -> %v", err)
	}
	// else {
	// 	defer os.RemoveAll(tmpDirPath)
	// }

	if err := os.MkdirAll(filepath.Join(tmpDirPath, "input", "{{Name}}", "{{Date}}"), 0744); err != nil {
		t.Fatalf("os.MkdirAll should have created template directories -> got error %v", err)
	}

	inputDir, outputDir := filepath.Join(tmpDirPath, "input", "{{Name}}"), filepath.Join(tmpDirPath, "output")

	args := []string{inputDir, outputDir}

	cmd.Use.Run(cmd.Use, args)
}
