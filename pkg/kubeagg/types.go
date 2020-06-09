package kubeagg

const (
	// kubectl executable
	kubectl = "kubectl"
)

// GlobalConfig global command line parameters values
type GlobalConfig struct {
	LogLevel string
}

// Config set command command line parameters values
type Config struct {
	Output         string
	Namespace      string
	ObjectType     string
	Contexts       []string
	ContextPattern string
}

// List of items in context
type List struct {
	Items   []Object
	Context string
}

// Object any object in the context
type Object struct {
	Kind     string
	Metadata Metadata
	Status   Status
}

// Status of the object
type Status struct {
	Phase string
}

// Metadata of the object
type Metadata struct {
	Name string
}

// AllObjects aggregated list of all contexts
type AllObjects struct {
	Lists []List
	Type  string
}
