package http

type (
	Handler map[ServiceName]Dependency

	Dependency struct {
		Sits Situations `yaml:"situations"`
	}
)
