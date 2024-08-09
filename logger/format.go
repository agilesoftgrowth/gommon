package logger

const (
	Json Format = iota
	Text
)

type Format int

func (f Format) String() string {
	switch f {
	case Json:
		return "json"
	default:
		return "text"
	}
}
