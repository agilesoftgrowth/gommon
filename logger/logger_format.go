package logger

import "strings"

const (
	FormatJson LoggerFormat = iota
	FormatText
)

type LoggerFormat int

func (f LoggerFormat) String() string {
	switch f {
	case FormatJson:
		return "json"
	default:
		return "text"
	}
}

func Format(s string) LoggerFormat {
	switch strings.ToLower(s) {
	case "json":
		return FormatJson
	default:
		return FormatText
	}
}
