package log

import "strings"

type Level int

const (
	ALERT = iota
	ERROR
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
	case "alert":
		return ALERT
	default:
		return DEBUG
	}
}

var levelPrefixes = []string{
	"ALERT",
	"ERROR",
	"WARNING",
	"INFO",
	"DEBUG",
}
