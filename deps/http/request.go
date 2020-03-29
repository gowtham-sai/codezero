package http

type request struct {
	Method  method `yaml:"method"`
	Path    string `yaml:"path"`
	Query   query  `yaml:"query"`
	Headers header `yaml:"headers"`
}
