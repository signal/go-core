package log

import (
	"fmt"
	"log"
)

type StderrLogger struct {
	level Level
}

func NewStderrLogger(l Level) *StderrLogger {
	return &StderrLogger{level: l}
}

func (s *StderrLogger) Debug(vals ...interface{}) {
	s.output(DEBUG, fmt.Sprint(vals...))
}

func (s *StderrLogger) Debugf(msg string, vals ...interface{}) {
	s.Debug(fmt.Sprintf(msg, vals...))
}

func (s *StderrLogger) Info(vals ...interface{}) {
	s.output(INFO, fmt.Sprint(vals...))
}

func (s *StderrLogger) Infof(msg string, vals ...interface{}) {
	s.Info(fmt.Sprintf(msg, vals...))
}

func (s *StderrLogger) Warn(vals ...interface{}) {
	s.output(WARN, fmt.Sprint(vals...))
}

func (s *StderrLogger) Warnf(msg string, vals ...interface{}) {
	s.Warn(fmt.Sprintf(msg, vals...))
}

func (s *StderrLogger) Error(vals ...interface{}) {
	s.output(ERROR, fmt.Sprint(vals...))
}

func (s *StderrLogger) Errorf(msg string, vals ...interface{}) {
	s.Error(fmt.Sprintf(msg, vals...))
}

func (s *StderrLogger) Fatal(vals ...interface{}) {
	log.Fatalln("FATAL", fmt.Sprint(vals...))
}

func (s *StderrLogger) Fatalf(msg string, vals ...interface{}) {
	s.Fatal(fmt.Sprintf(msg, vals...))
}

func (s *StderrLogger) LogDebug() bool {
	return s.level >= DEBUG
}

func (s *StderrLogger) LogWarn() bool {
	return s.level >= WARN
}

func (s *StderrLogger) LogInfo() bool {
	return s.level >= INFO
}

func (s *StderrLogger) LogError() bool {
	return s.level >= ERROR
}

func (s *StderrLogger) LogLevel(level Level) bool {
	return s.level >= level
}

func (s *StderrLogger) output(level Level, vals ...interface{}) {
	if level <= s.level {
		log.Println(levelPrefixes[level], fmt.Sprint(vals...))
	}
}
