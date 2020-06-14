package kubeagg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

// Output in provided format
func (c *Contexts) Output(outputType string) {
	switch outputType {
	case "json":
		c.OutputJSON()
	case "table":
		c.OutputTable()
	// case "wide":
	// 	objects.PrintWide()
	default:
		sugar.Warnf("Output type %v is not supported, supported values is json, table, wide", outputType)
	}
}

// PrintTable in table format
func (c *Contexts) OutputTable() {
	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	if !globalConfigVar.NoHeaders {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t", "CONTEXT", "NAMESPACE", "TYPE", "NAME")
	}

	// Namespaced
	for i := 0; i < len(c.Contexts); i++ {
		for j := 0; j < len(c.Contexts[i].Namespaces); j++ {
			for _, object := range c.Contexts[i].Namespaces[j].Objects.Items {
				fmt.Fprintf(w, "\n%v\t%v\t%v\t%v", c.Contexts[i].Name, c.Contexts[i].Namespaces[j].Name, object.Kind, object.Metadata.Name)
			}
		}
	}
	// NonNamespaced
	for i := 0; i < len(c.Contexts); i++ {
		for _, object := range c.Contexts[i].NonNamespaced.Items {
			fmt.Fprintf(w, "\n%v\t%v\t%v\t%v", c.Contexts[i].Name, "n\\a", object.Kind, object.Metadata.Name)
		}
	}

	fmt.Fprintln(w, "")

}

func (c *Contexts) OutputJSON() {
	json, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}
