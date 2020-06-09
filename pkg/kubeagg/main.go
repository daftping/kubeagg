package kubeagg

import (
	"encoding/json"
	"os/exec"
	"sync"
)

// Run package entrypoint
func Run() {

	// Generate final contexts lists
	contexts := GetContexts()
	sugar.Debugw(
		"Contexts to precess",
		"contexts", contexts,
	)

	// Async context processing
	var (
		wg        sync.WaitGroup
		allObject AllObjects
	)

	for _, context := range contexts {
		wg.Add(1)
		go func(context string) {
			allObject.Lists = append(allObject.Lists, GetFromContext(context))
			wg.Done()
		}(context)

	}
	wg.Wait()

	// Sync
	// var allObject AllObjects

	// for _, context := range contexts {
	// 	allObject.Lists = append(allObject.Lists, GetFromContext(context))
	// }

	// Output results
	allObject.Output(getConfigVar.Output)
}

// GetFromContext get objects from individual context
func GetFromContext(context string) (ctxObjects List) {
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

	return
}
