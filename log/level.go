package log

import "strings"

type Level int

const (
	ERROR = iota
	WARN
	INFO
	DEBUG
)

func parseLevel(lvl string) Level {
	switch strings.ToLower(lvl) {
	case "error":
		return ERROR
	case "warn", "warning":
		return WARN
	case "info":
		return INFO
	default:
		return DEBUG
	}
}

var levelPrefixes = []string{
	"ERROR",
	"WARNING",
	"INFO",
	"DEBUG",
}
