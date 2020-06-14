package kubeagg

import "regexp"

// GetNamespaces Stub for namespace parsing logic
func (c *Contexts) GetNamespaces() {
	// Check if namespaces is provided, empty slice by default.
	// --namespaces has precedence over --namespace-pattern
	// --namespace-pattern is ignored
	// if --namespaces is provided, fill in Contexts struct with namespaces
	if len(globalConfigVar.Namespaces) > 0 {
		//TODO probably should add * as prefix and suffix for those who doesn't like regexp like me ))
		for i := 0; i < len(c.Contexts); i++ {
			for _, ns := range globalConfigVar.Namespaces {
				c.Contexts[i].Namespaces = append(
					c.Contexts[i].Namespaces,
					Namespace{
						Name: ns,
					},
				)
			}
		}
		sugar.Debugw(
			"Aggregated uniq list of namespaces to process",
			"namespaces", c.GetNamespacesNames(),
		)
		return
	}

	// If namespace pattern is used
	c.PopulateNamespacesByPatternAsync()
	sugar.Debugw(
		"Aggregated uniq list of namespaces to process",
		"namespaces", c.GetNamespacesNames(),
	)
}

//GetNamespacesNames helped function for debug output
func (c *Contexts) GetNamespacesNames() (s []string) {
	for _, context := range c.Contexts {
		for _, namespace := range context.Namespaces {
			s = append(s, namespace.Name)
		}
	}

	return unique(s)
}

// GetNamespacesByPattern filter out namespaces by pattern
func GetContextNamespacesByPattern(pattern string, allNamespaces List) (filteredNamespaces []Namespace) {
	var desiredNamespaces = regexp.MustCompile(pattern)
	for _, namespace := range allNamespaces.Items {
		if desiredNamespaces.MatchString(namespace.Metadata.Name) {
			filteredNamespaces = append(
				filteredNamespaces,
				Namespace{
					Name: namespace.Metadata.Name,
				},
			)
		}
	}
	return
}

func (c *Contexts) PopulateNamespacesByPatternAsync() {
	sugar.Debugf("Namespace pattern is used: %v", globalConfigVar.NamespacePattern)
	sugar.Debug("Getting namespaces from contexts")

	for i := 0; i < len(c.Contexts); i++ {
		c.WaitGroup.Add(1)
		go func(i int) {
			Namespaces := GetObjectsFromContext(
				GetKubectlArgs(c.Contexts[i].Name, "", "namespace"),
			)
			// Assign filtered namespaces
			c.Contexts[i].Namespaces = GetContextNamespacesByPattern(globalConfigVar.NamespacePattern, Namespaces)
			c.WaitGroup.Done()
		}(i)
	}
	sugar.Debug("Waiting for all requests to complete")
	c.WaitGroup.Wait()
	return
}
