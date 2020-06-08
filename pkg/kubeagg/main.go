package kubeagg

import (
	"encoding/json"
	"os/exec"
)

func Run() {
	contexts := GetContexts()
	sugar.Debugw(
		"Contexts to precess",
		"contexts", contexts,
	)

	var allObject AllObjects

	allObject.Type = getConfigVar.ObjectType

	// TODO run async
	for _, context := range contexts {
		var ctxObjects List

		args := []string{
			"get",
			getConfigVar.ObjectType,
			"--output=json",
			"--context", context,
			"--namespace", getConfigVar.Namespace,
		}

		sugar.Debugf("Running: %v %v", kubectl, args)

		out, err := exec.Command(kubectl, args...).Output()
		if err != nil {
			sugar.Fatal(err)
		}

		errJSON := json.Unmarshal(out, &ctxObjects)
		if errJSON != nil {
			sugar.Fatal(errJSON)
		}

		ctxObjects.Context = context

		allObject.Lists = append(allObject.Lists, ctxObjects)
	}

	switch getConfigVar.Output {
	case "json":
		allObject.PrintJSON()
	case "table":
		allObject.PrintTable()
	case "wide":
		allObject.PrintWide()
	default:
		sugar.Warnf("Output type %v is not supported, supported values is json, table, wide", getConfigVar.Output)
	}

}
