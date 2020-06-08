package kubeagg

import (
	"encoding/json"
	"log"
	"os/exec"
	"regexp"
)

// GetContexts converts provided as parameter --context=... contexts and patters
// as list of contexts
func GetContexts() []string {
	//Check if contexts is provided, empty slice by default.
	// --contexts has precedence over --context-pattern
	// --context-pattern is ignored
	if len(getConfigVar.Contexts) > 0 {
		//Support * as context.
		//Does it make sense?
		//TODO probably should add * as prefix and suffix for those who doesn't like regexp like me ))
		if isWildcard() {
			return GetAllContexts()
		}
		return getConfigVar.Contexts
	}

	//If context pattern is used
	return GetContextsByPattern(getConfigVar.ContextPattern)
}

// GetContextsByPattern converts context pattern regexp into slice of contexts
func GetContextsByPattern(pattern string) (contexts []string) {
	var desiredContexts = regexp.MustCompile(pattern)
	for _, context := range GetAllContexts() {
		if desiredContexts.MatchString(context) {
			contexts = append(contexts, context)
		}
	}
	return
}

// Check if any of provided contexts is *
func isWildcard() bool {
	for _, context := range getConfigVar.Contexts {
		if context == "*" {
			return true
		}
	}
	return false
}

// GetAllContexts gets all context in ~/.kube/config
func GetAllContexts() []string {
	type Context struct {
		Name string
	}
	type KubeConfig struct {
		Contexts []Context
	}
	var kubeconfig KubeConfig
	var contexts []string

	out, errExec := exec.Command(
		kubectl,
		"config",
		"view",
		"--output=json",
	).Output()
	if errExec != nil {
		log.Fatal(errExec)
	}

	errJSON := json.Unmarshal(out, &kubeconfig)
	if errJSON != nil {
		log.Fatal(errJSON)
	}

	for _, context := range kubeconfig.Contexts {
		contexts = append(contexts, context.Name)
	}

	return contexts
}
