package logrus

import "github.com/sirupsen/logrus"

type Logger struct {
	Entity *logrus.Entry
}

func New() *Logger {
	log := logrus.New()
	log.Level = logrus.DebugLevel

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
	})

	return &Logger{
		Entity: logrus.NewEntry(log),
	}
}

func (l Logger) Debug(v ...interface{}) {
	l.Entity.Debug(v...)
}
func (l Logger) Debugf(format string, v ...interface{}) {
	l.Entity.Debugf(format, v...)
}
func (l Logger) JDebug(fields map[string]interface{}, v ...interface{}) {
	l.Entity.WithFields(fields).Debug(v...)
}
func (l Logger) JDebugf(fields map[string]interface{}, format string, v ...interface{}) {
	l.Entity.WithFields(fields).Debugf(format, v...)
}

func (l Logger) Info(v ...interface{}) {
	l.Entity.Info(v...)
}
func (l Logger) Infof(format string, v ...interface{}) {
	l.Entity.Infof(format, v...)
}
func (l Logger) JInfo(fields map[string]interface{}, v ...interface{}) {
	l.Entity.WithFields(fields).Info(v...)
}
func (l Logger) JInfof(fields map[string]interface{}, format string, v ...interface{}) {
	l.Entity.WithFields(fields).Infof(format, v...)
}

func (l Logger) Warn(v ...interface{}) {
	l.Entity.Warn(v)
}
func (l Logger) Warnf(format string, v ...interface{}) {
	l.Entity.Warnf(format, v...)
}
func (l Logger) JWarn(fields map[string]interface{}, v ...interface{}) {
	l.Entity.WithFields(fields).Warn(v...)
}
func (l Logger) JWarnf(fields map[string]interface{}, format string, v ...interface{}) {
	l.Entity.WithFields(fields).Warnf(format, v...)
}

func (l Logger) Error(v ...interface{}) {
	l.Entity.Error(v...)
}
func (l Logger) Errorf(format string, v ...interface{}) {
	l.Entity.Errorf(format, v...)
}
func (l Logger) JError(fields map[string]interface{}, v ...interface{}) {
	l.Entity.WithFields(fields).Error(v...)
}
func (l Logger) JErrorf(fields map[string]interface{}, format string, v ...interface{}) {
	l.Entity.WithFields(fields).Errorf(format, v...)
}

func (l Logger) Fatal(v ...interface{}) {
	l.Entity.Fatal(v...)
}
func (l Logger) Fatalf(format string, v ...interface{}) {
	l.Entity.Fatalf(format, v...)
}
func (l Logger) JFatal(fields map[string]interface{}, v ...interface{}) {
	l.Entity.WithFields(fields).Fatal(v...)
}
func (l Logger) JFatalf(fields map[string]interface{}, format string, v ...interface{}) {
	l.Entity.WithFields(fields).Fatalf(format, v...)
}
func (l Logger) WithFields(fields map[string]interface{}) Logger {
	l.Entity = l.Entity.WithFields(fields)
	return l
}
