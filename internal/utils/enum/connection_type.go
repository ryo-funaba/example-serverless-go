package enum

type ConnectionType int

const (
	Primary ConnectionType = iota
	Secondary
)

func (t ConnectionType) String() string {
	switch t {
	case Primary:
		return "primary"
	case Secondary:
		return "secondary"
	default:
		return ""
	}
}
