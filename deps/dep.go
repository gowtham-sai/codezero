package deps

var (
	handlers = map[Type]Handler{}
)

type (
	Type string
	Spec string

	Handler interface {
		StartSituation() error
		ParseSpec(spec Spec) error
	}
)

func RegisterHandler(t Type, h Handler) error {
	return nil
}
