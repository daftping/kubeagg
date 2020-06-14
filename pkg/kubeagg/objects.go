package kubeagg

func (c *Contexts) PopulateNamespacedObjectsAsync() {
	sugar.Debugf("Getting object: %v", getConfigVar.ObjectType)
	for i := 0; i < len(c.Contexts); i++ {
		c.WaitGroup.Add(1)
		go func(i int) {
			for j := 0; j < len(c.Contexts[i].Namespaces); j++ {
				c.Contexts[i].Namespaces[j].Objects = GetObjectsFromContext(
					GetKubectlArgs(c.Contexts[i].Name, c.Contexts[i].Namespaces[j].Name, getConfigVar.ObjectType),
				)
			}
			c.WaitGroup.Done()
		}(i)
	}
	sugar.Debug("Waiting for all requests to complete")
	c.WaitGroup.Wait()
	return
}

func (c *Contexts) PopulateNonNamespacedObjectsAsync() {
	sugar.Debugf("Getting object: %v", getConfigVar.ObjectType)
	for i := 0; i < len(c.Contexts); i++ {
		c.WaitGroup.Add(1)
		go func(i int) {
			c.Contexts[i].NonNamespaced = GetObjectsFromContext(
				GetKubectlArgs(c.Contexts[i].Name, "", getConfigVar.ObjectType),
			)
			c.WaitGroup.Done()
		}(i)
	}
	sugar.Debug("Waiting for all requests to complete")
	c.WaitGroup.Wait()
	return
}
