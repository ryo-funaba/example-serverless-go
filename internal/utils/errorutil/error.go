package errorutil

type Code uint

const (
	_ Code = iota
	OK
	Unknown
)

func (c Code) Name() string {
	switch c {
	case OK:
		return "ok"
	case Unknown:
		return "unknown"
	default:
		return "unknown"
	}
}

func (c Code) Number() int {
	switch c {
	case OK:
		return 200
	case Unknown:
		return 500
	default:
		return 500
	}
}
