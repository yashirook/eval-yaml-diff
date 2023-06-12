package gateway

import (
	"eval-yaml-diff/internal/domain"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type PrintGateway struct{}

func (pg PrintGateway) Print(diffs domain.DiffList) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"PATH", "CHANGE_TYPE", "EVAL_RESULT"})

	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetBorder(false)

	for _, diff := range diffs {
		var result string = "DENIED"
		if diff.Allowed {
			result = "ALLOWED"
		}
		row := []string{diff.Path, strings.ToUpper(string(diff.ChangeType)), result}
		table.Append(row)
	}

	table.Render()
}
