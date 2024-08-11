package logger

const (
	FormatJson Format = iota
	FormatText
)

type Format int

func (f Format) String() string {
	switch f {
	case FormatJson:
		return "json"
	default:
		return "text"
	}
}
