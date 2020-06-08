package kubeagg

var globalConfigVar = GlobalConfig{}
var getConfigVar = Config{}

// SetGlobalConfig stores passed config in global package variables
func SetGlobalConfig(global GlobalConfig) {
	globalConfigVar = global
	initZapLog()
}

// SetConfig stores passed config in global package variables
func SetConfig(config Config) {
	getConfigVar = config
}
