package kubeagg

import (
	"encoding/json"
	"log"
	"os/exec"
)

type List struct {
	Items   []Object
	Context string
}

type Object struct {
	Metadata Metadata
}

type Metadata struct {
	Name string
}

type AllObjects struct {
	Lists []List
}

func Run(config Config) {
	// TODO Make this as customizable
	contexts := []string{"docker-desktop", "test-mc-e"}

	var allObject AllObjects
	var ctxObjects List

	// TODO run async
	for _, context := range contexts {
		// TODO make object type as parameters
		// TODO make output as parameter
		out, err := exec.Command("kubectl", "get", "ns", "-o=json", "--context", context).Output()
		if err != nil {
			log.Fatal(err)
		}

		err2 := json.Unmarshal(out, &ctxObjects)
		if err2 != nil {
			log.Fatal(err2)
		}

		ctxObjects.Context = context
		// fmt.Printf("%+v\n", ctxObjects)

		allObject.Lists = append(allObject.Lists, ctxObjects)
	}

	switch config.Output {
	case "json":
		allObject.PrintJSON()
	case "table":
		allObject.PrintTable()
	}

}
