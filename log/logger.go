package log

type Logger interface {
	Debug(vals ...interface{})
	Debugf(msg string, vals ...interface{})

	Info(vals ...interface{})
	Infof(msg string, vals ...interface{})

	Warn(vals ...interface{})
	Warnf(msg string, vals ...interface{})

	Error(vals ...interface{})
	Errorf(msg string, vals ...interface{})

	Fatal(vals ...interface{})
	Fatalf(msg string, vals ...interface{})

	Alert(vals ...interface{})
	Alertf(msg string, vals ...interface{})

	LogDebug() bool
	LogWarn() bool
	LogInfo() bool
	LogError() bool
}

var logger Logger = NewStderrLogger(DEBUG)

func SetLogger(l Logger) {
	logger = l
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Alert(args ...interface{}) {
	logger.Alert(args...)
}

func Alertf(format string, args ...interface{}) {
	logger.Alertf(format, args...)
}

func LogDebug() bool {
	return logger.LogDebug()
}

func LogWarn() bool {
	return logger.LogWarn()
}

func LogInfo() bool {
	return logger.LogInfo()
}

func LogError() bool {
	return logger.LogError()
}
