package kubeagg

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

func Run(config Config) {
	// TODO Add debug logging
	fmt.Println(GetContexts(config))
	contexts := GetContexts(config)

	var allObject AllObjects

	allObject.Type = config.ObjectType

	// TODO run async
	for _, context := range contexts {
		var ctxObjects List
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
		log.Printf("Output type %v is not supported, supported values is json, table, wide", config.Output)
	}

}
