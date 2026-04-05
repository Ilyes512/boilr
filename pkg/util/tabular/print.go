package tabular

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

// Print writes the given header and data in a tabular format to stdout.
func Print(header []string, data [][]string) error {
	symbols := tw.NewSymbolCustom("boilr").
		WithCenter(color.GreenString("+")).
		WithColumn(color.GreenString("|")).
		WithRow(color.GreenString("-"))

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{Symbols: symbols})),
	)

	headerAny := make([]any, len(header))
	for i, h := range header {
		headerAny[i] = h
	}
	table.Header(headerAny...)

	for _, datum := range data {
		datum[0] = color.RedString(datum[0])

		if strings.HasPrefix(datum[1], "local:") {
			datum[1] = color.BlueString(datum[1])
		} else {
			datum[1] = color.YellowString(datum[1])
		}

		if err := table.Append(datum); err != nil {
			return err
		}
	}

	if len(data) == 0 {
		if err := table.Append([]string{"", "", ""}); err != nil {
			return err
		}
	}

	return table.Render()
}
