package template

import (
	"text/template"
	"time"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/group/all"
)

var (
	// Options contain the default options for the template execution.
	Options = []string{
		// TODO ignore a field if no value is found instead of writing <no value>
		"missingkey=invalid",
	}
)

// buildFuncMap creates a fresh template.FuncMap with all sprout built-in functions
// and the custom boilr functions registered via boilrRegistry.
func buildFuncMap() template.FuncMap {
	h := sprout.New(
		sprout.WithGroups(all.RegistryGroup()),
		sprout.WithRegistries(newBoilrRegistry()),
	)
	return h.Build()
}

// CurrentTimeInFmt returns the current time in the given format.
// See time.Time.Format for more details on the format string.
func CurrentTimeInFmt(fmt string) string {
	t := time.Now()

	return t.Format(fmt)
}
