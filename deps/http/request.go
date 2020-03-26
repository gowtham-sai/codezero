package http

type Request struct {
	Method  Method `yaml:"method"`
	Path    string `yaml:"path"`
	Query   Query  `yaml:"query"`
	Headers Header `yaml:"headers"`
}
