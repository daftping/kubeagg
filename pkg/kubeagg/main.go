package kubeagg

import (
	"encoding/json"
	"log"
	"os/exec"
)

const (
	kubectl = "kubectl"
)

func Run(config Config) {
	// TODO Make this as customizable
	contexts := []string{"docker-desktop", "test-mc-e"}

	var allObject AllObjects

	allObject.Type = config.ObjectType

	// TODO run async
	for _, context := range contexts {
		var ctxObjects List
		// TODO make object type as parameters
		// TODO make output as parameter
		// TODO debug command output
		out, err := exec.Command(
			kubectl,
			"get",
			config.ObjectType,
			"--output=json",
			"--context", context,
			"--namespace", config.Namespace,
		).Output()
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
	case "wide":
		allObject.PrintWide()
	default:
		log.Printf("Output type %v is not supported, supported values is json, table", config.Output)
	}

}
