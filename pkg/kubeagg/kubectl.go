package kubeagg

import (
	"encoding/json"
	"os/exec"
)

const (
	// kubectl executable
	kubectl = "kubectl"
)

// GetObjectsFromContext get objects from individual context
func GetObjectsFromContext(args []string) (list List) {
	sugar.Debugf("Running: %v %v", kubectl, args)

	out, err := exec.Command(kubectl, args...).Output()
	if err != nil {
		sugar.Fatal(err)
	}

	errJSON := json.Unmarshal(out, &list)
	if errJSON != nil {
		sugar.Fatal(errJSON)
	}
	return
}

// GetArgs get arguments for kubectl
func GetKubectlArgs(context string, namespace string, objectType string) (args []string) {
	args = []string{
		"get",
		objectType,
		"--output=json",
		"--context", context,
	}

	if namespace != "" {
		args = append(args, "--namespace="+namespace)
	}
	return
}
