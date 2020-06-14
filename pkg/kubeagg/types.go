package kubeagg

import (
	"sync"
	"time"
)

// GetNonNamespacedObjects get non namespaced objects
// Should remove this hardcode and make decision based on cluster objects types
// # In a namespace
// kubectl api-resources --namespaced=true
// # Not in a namespace
// kubectl api-resources --namespaced=false
func GetNonNamespacedObjects() []string {
	return []string{
		"componentstatuses", "cs",
		"namespaces", "ns",
		"node", "nodes", "no",
		"persistentvolumes", "pv",
		"mutatingwebhookconfigurations",
		"validatingwebhookconfigurations",
		"customresourcedefinitions", "crd", "crds",
		"apiservices",
		"tokenreviews",
		"selfsubjectaccessreviews",
		"selfsubjectrulesreviews",
		"subjectaccessreviews",
		"certificatesigningrequests", "csr",
		"compositecontrollers", "cc", "cctl",
		"decoratorcontrollers", "dec", "decorators",
		"runtimeclasses",
		"podsecuritypolicies", "psp",
		"clusterrolebindings",
		"clusterroles",
		"priorityclasses", "pc",
		"csidrivers",
		"csinodes",
		"storageclasses", "sc",
		"volumeattachments",
	}
}

// Configs passed from commandline //

// GlobalConfig global command line parameters values
type GlobalConfig struct {
	LogLevel         string
	Output           string
	NoHeaders        bool
	Contexts         []string
	ContextPattern   string
	Namespaces       []string
	NamespacePattern string
}

// Config set command command line parameters values
type Config struct {
	ObjectType string
}

// Kubernetes JSON structure

// Object Kubernetes object root structure
type Object struct {
	Kind     string
	Metadata Metadata
	Status   Status
}

//List of kubernetes objects
type List struct {
	Items []Object
}

// Status kubernetes objects status
type Status struct {
	Phase string
}

// Metadata kubernetes object metadata
type Metadata struct {
	Name              string
	Namespace         string
	CreationTimestamp time.Time
}

// Aggregated structure

type Contexts struct {
	Contexts  []Context
	WaitGroup sync.WaitGroup
}

type Context struct {
	Name          string
	Namespaces    []Namespace
	NonNamespaced List
}

type Namespace struct {
	Name    string
	Objects List
}
