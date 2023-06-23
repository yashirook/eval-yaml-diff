package gateway

import (
	"eval-yaml-diff/internal/domain"
	"fmt"
	"os"
	"text/tabwriter"
)

type PrintGateway struct{}

func (pg PrintGateway) Print(diffs domain.DiffList) error {
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', tabwriter.AlignRight)
	fmt.Fprintln(tw, "PATH\tCHANGE_TYPE\tRESULT\t")

	for _, diff := range diffs {
		var result string = "DENIED"
		if diff.Allowed {
			result = "ALLOWED"
		}
		fmt.Fprintf(tw, "%s\t%s\t%s\n", diff.Path, diff.ChangeType, result)
	}

	if err := tw.Flush(); err != nil {
		return err
	}
	return nil
}
