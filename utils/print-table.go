package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// PrintTable returns a formatted table to be displayed
func PrintTable(aa map[string]string) {
	var data = map[string]string{
		"bra": "Brazil",
		"usa": "United States of America",
	}
	// minwidth, tabwidth, padding, padchar, flags
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 5, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t", "command", "country")
	fmt.Fprintf(w, "\n %s\t%s\t", "-------", "-------")

	for key, value := range data {
		fmt.Fprintf(w, "\n %v\t%v\t\n", key, value)
	}
}
