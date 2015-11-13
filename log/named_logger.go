package log

import "fmt"

type NamedLogger struct {
	name   string
	logger Logger
}

func NewNamedLogger(name string, logger Logger) *NamedLogger {
	return &NamedLogger{"[" + name + "] ", logger}
}

func (s *NamedLogger) Debug(vals ...interface{}) {
	s.logger.Debug(s.name + fmt.Sprint(vals...))
}

func (s *NamedLogger) Debugf(msg string, vals ...interface{}) {
	s.logger.Debugf(s.name+msg, vals...)
}

func (s *NamedLogger) Info(vals ...interface{}) {
	s.logger.Info(s.name + fmt.Sprint(vals...))
}

func (s *NamedLogger) Infof(msg string, vals ...interface{}) {
	s.logger.Infof(s.name+msg, vals...)
}

func (s *NamedLogger) Warn(vals ...interface{}) {
	s.logger.Warn(s.name + fmt.Sprint(vals...))
}

func (s *NamedLogger) Warnf(msg string, vals ...interface{}) {
	s.logger.Warnf(s.name+msg, vals...)
}

func (s *NamedLogger) Error(vals ...interface{}) {
	s.logger.Error(s.name + fmt.Sprint(vals...))
}

func (s *NamedLogger) Errorf(msg string, vals ...interface{}) {
	s.logger.Errorf(s.name+msg, vals...)
}

func (s *NamedLogger) Fatal(vals ...interface{}) {
	s.logger.Fatal(s.name + fmt.Sprint(vals...))
}

func (s *NamedLogger) Fatalf(msg string, vals ...interface{}) {
	s.logger.Fatalf(s.name+msg, vals...)
}

func (s *NamedLogger) LogDebug() bool {
	return s.logger.LogDebug()
}

func (s *NamedLogger) LogWarn() bool {
	return s.logger.LogWarn()
}

func (s *NamedLogger) LogInfo() bool {
	return s.logger.LogInfo()
}

func (s *NamedLogger) LogError() bool {
	return s.logger.LogError()
}
