package kubeagg

import "sync"

// New return Contexts object
func New() *Contexts {
	contexts := Contexts{}
	contexts.WaitGroup = sync.WaitGroup{}
	return &contexts
}

// Run package entrypoint
func Run() {
	// Get reference to Contexts object
	c := New()

	// Populate contexts
	c.GetContexts()

	// Handle namespaced and non namespaced object differently
	if contains(GetNonNamespacedObjects(), getConfigVar.ObjectType) {
		c.PopulateNonNamespacedObjectsAsync()

	} else {
		c.GetNamespaces()
		c.PopulateNamespacedObjectsAsync()
	}

	// Output in provided format
	c.Output(globalConfigVar.Output)
}
