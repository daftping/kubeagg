package kubeagg

import (
	"encoding/json"
	"os/exec"
	"regexp"
)

// GetContexts make a decision whether use --contexts or --context-pattern
func (c *Contexts) GetContexts() {
	// Check if contexts is provided, empty slice by default.
	// --contexts has precedence over --context-pattern
	// --context-pattern is ignored
	if len(globalConfigVar.Contexts) > 0 {
		//Support * as context.
		//Does it make sense?
		//TODO probably should add * as prefix and suffix for those who doesn't like regexp like me ))
		if isWildcard() {
			c.Append(GetAllContexts())
			return
		}
		//Add provided contexts as is
		c.Append(globalConfigVar.Contexts)
		sugar.Debugw(
			"Contexts to process",
			"contexts", c.GetContextsNames(),
		)
		return
	}

	// If context pattern is used resolve pattern and add to Contexts
	c.Append(GetContextsByPattern(globalConfigVar.ContextPattern))
	sugar.Debugw(
		"Contexts to process",
		"contexts", c.GetContextsNames(),
	)
	return
}

// Append convert []string to []Context and append to Contexts
func (c *Contexts) Append(contexts []string) {
	for _, context := range contexts {
		c.Contexts = append(c.Contexts, Context{
			Name: context,
		})
	}
}

func (c *Contexts) GetContextsNames() (s []string) {
	for _, context := range c.Contexts {
		s = append(s, context.Name)
	}
	return
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
	for _, context := range globalConfigVar.Contexts {
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
		sugar.Fatal(errExec)
	}

	errJSON := json.Unmarshal(out, &kubeconfig)
	if errJSON != nil {
		sugar.Fatal(errJSON)
	}

	for _, context := range kubeconfig.Contexts {
		contexts = append(contexts, context.Name)
	}

	return contexts
}
