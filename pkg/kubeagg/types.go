package kubeagg

const (
	kubectl = "kubectl"
)

type GlobalConfig struct {
	LogLevel string
}

type Config struct {
	Output         string
	Namespace      string
	ObjectType     string
	Contexts       []string
	ContextPattern string
}

type List struct {
	Items   []Object
	Context string
}

type Object struct {
	Metadata Metadata
}

type Metadata struct {
	Name string
}

type AllObjects struct {
	Lists []List
	Type  string
}
