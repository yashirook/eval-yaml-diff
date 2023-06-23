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
	if _, err := fmt.Fprintln(tw, "PATH\tCHANGE_TYPE\tRESULT\t"); err != nil {
		return err
	}

	for _, diff := range diffs {
		var result = "DENIED"
		if diff.Allowed {
			result = "ALLOWED"
		}
		if _, err := fmt.Fprintf(tw, "%s\t%s\t%s\n", diff.Path, diff.ChangeType, result); err != nil {
			return err
		}

	}

	if err := tw.Flush(); err != nil {
		return err
	}
	return nil
}
