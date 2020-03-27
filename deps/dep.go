package deps

var (
	handlers = map[Type]Handler{}
)

type (
	Type string
	Spec string

	Handler interface {
		StartSituation(spec Spec) error
		StopSituation(spec Spec) error
		ParseSpec(spec Spec) error
	}
)

func RegisterHandler(t Type, h Handler) error {
	return nil
}
