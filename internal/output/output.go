package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"gopkg.in/yaml.v3"
)

// Print writes data in the specified format to stdout.
// For "table" format, columns defines the header and data is expected to be [][]string.
// For "json" and "yaml" formats, data can be any serializable value.
func Print(format string, data interface{}, columns []string) error {
	return Fprint(os.Stdout, format, data, columns)
}

// Fprint writes data in the specified format to the given writer.
func Fprint(w io.Writer, format string, data interface{}, columns []string) error {
	switch format {
	case "json":
		return printJSON(w, data)
	case "yaml":
		return printYAML(w, data)
	case "table", "":
		return printTable(w, data, columns)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
}

func printJSON(w io.Writer, data interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func printYAML(w io.Writer, data interface{}) error {
	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	defer enc.Close()
	return enc.Encode(data)
}

func printTable(w io.Writer, data interface{}, columns []string) error {
	rows, ok := data.([][]string)
	if !ok {
		return fmt.Errorf("table format requires [][]string data")
	}

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	if len(columns) > 0 {
		fmt.Fprintln(tw, strings.Join(columns, "\t"))
	}

	for _, row := range rows {
		fmt.Fprintln(tw, strings.Join(row, "\t"))
	}

	return tw.Flush()
}
