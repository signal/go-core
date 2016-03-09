package log

import (
	"fmt"

	"gopkg.in/bugsnag/bugsnag-go.v1"
)

type BugsnagLogger struct {
	name   string
	logger Logger
	config *BugsnagConfig
}

type BugsnagConfig struct {
	env                string
	version            string
	errorServiceAPIKey string
	reportable         bool
}

func NewBugsnagConfig(version, env, apiKey string) *BugsnagConfig {
	reportable := false
	if env == "" || apiKey == "" {
		Warn("Bugsnag error reporting is disabled (missing environment name or API key)")
	} else {
		bugsnag.Configure(bugsnag.Configuration{
			APIKey:              apiKey,
			ReleaseStage:        env,
			AppVersion:          version,
			NotifyReleaseStages: []string{"production"},
			PanicHandler:        func() {},
		})
		reportable = true
	}
	return &BugsnagConfig{env, version, apiKey, reportable}
}

func NewBugsnagLogger(name, build, env, logLevel, apiKey string) *BugsnagLogger {
	return &BugsnagLogger{"[" + name + "] ", NewStderrLogger(parseLevel(logLevel)),
		NewBugsnagConfig(build, env, apiKey)}
}

func (s *BugsnagLogger) Debug(vals ...interface{}) {
	s.logger.Debug(s.name + fmt.Sprint(vals...))
}

func (s *BugsnagLogger) Debugf(msg string, vals ...interface{}) {
	s.logger.Debugf(s.name+msg, vals...)
}

func (s *BugsnagLogger) Info(vals ...interface{}) {
	s.logger.Info(s.name + fmt.Sprint(vals...))
}

func (s *BugsnagLogger) Infof(msg string, vals ...interface{}) {
	s.logger.Infof(s.name+msg, vals...)
}

func (s *BugsnagLogger) Warn(vals ...interface{}) {
	s.logger.Warn(s.name + fmt.Sprint(vals...))
}

func (s *BugsnagLogger) Warnf(msg string, vals ...interface{}) {
	s.logger.Warnf(s.name+msg, vals...)
}

func (s *BugsnagLogger) Error(vals ...interface{}) {
	bugSnagVals := append(vals, bugsnag.ErrorClass{"Error"})
	s.notify(fmt.Errorf(fmt.Sprint(vals...)), bugSnagVals...)
	s.logger.Error(s.name + fmt.Sprint(vals...))
}

func (s *BugsnagLogger) Errorf(msg string, vals ...interface{}) {
	bugSnagVals := append(vals, bugsnag.ErrorClass{"Error"})
	s.notify(fmt.Errorf(msg, vals...), bugSnagVals...)
	s.logger.Errorf(s.name+msg, vals...)
}

func (s *BugsnagLogger) Fatal(vals ...interface{}) {
	// fatal errors need to be synchronous so the program doesn't exit before bugsnag is notified
	bugSnagVals := append(vals, bugsnag.Configuration{Synchronous: true})
	bugSnagVals = append(bugSnagVals, bugsnag.ErrorClass{"Fatal"})
	s.notify(fmt.Errorf(fmt.Sprint(vals...)), bugSnagVals...)
	s.logger.Fatal(s.name + fmt.Sprint(vals...))
}

func (s *BugsnagLogger) Fatalf(msg string, vals ...interface{}) {
	bugSnagVals := append(vals, bugsnag.Configuration{Synchronous: true})
	bugSnagVals = append(bugSnagVals, bugsnag.ErrorClass{"Fatal"})
	s.notify(fmt.Errorf(msg, vals...), bugSnagVals...)
	s.logger.Fatalf(s.name+msg, vals...)
}

func (s *BugsnagLogger) LogDebug() bool {
	return s.logger.LogDebug()
}

func (s *BugsnagLogger) LogWarn() bool {
	return s.logger.LogWarn()
}

func (s *BugsnagLogger) LogInfo() bool {
	return s.logger.LogInfo()
}

func (s *BugsnagLogger) LogError() bool {
	return s.logger.LogError()
}

func (s *BugsnagLogger) notify(err error, data ...interface{}) {
	if s.config.reportable {
		data = append(data, bugsnag.SeverityError)
		bugsnag.Notify(err, data...)
	}
}
