package kubeagg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

// Output in provided format
func (objects AllObjects) Output(output string) {
	switch output {
	case "json":
		objects.PrintJSON()
	case "table":
		objects.PrintTable()
	case "wide":
		objects.PrintWide()
	default:
		sugar.Warnf("Output type %v is not supported, supported values is json, table, wide", output)
	}
}

// PrintTable in table format
func (objects AllObjects) PrintTable() {
	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "%s\t%s\t", "CONTEXT", "NAME")

	for _, list := range objects.Lists {
		for _, item := range list.Items {
			fmt.Fprintf(w, "\n %v\t%v", list.Context, item.Metadata.Name)
		}
	}
	fmt.Fprintln(w, "")
}

// PrintWide same as PrintTable but with additional fields
func (objects AllObjects) PrintWide() {
	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t", "CONTEXT", "TYPE", "NAME", "STATUS")

	for _, list := range objects.Lists {
		for _, item := range list.Items {
			fmt.Fprintf(w, "\n %v\t%v\t%v\t%v", list.Context, item.Kind, item.Metadata.Name, item.Status.Phase)
		}
	}
	fmt.Fprintln(w, "")
}

// PrintJSON output in JSON format
func (objects AllObjects) PrintJSON() {
	json, err := json.MarshalIndent(objects, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}
