package logger

type Logger interface {
	JDebug(fields map[string]interface{}, v ...interface{})
	JDebugf(fields map[string]interface{}, format string, v ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	JInfo(fields map[string]interface{}, v ...interface{})
	JInfof(fields map[string]interface{}, format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	JWarn(fields map[string]interface{}, v ...interface{})
	JWarnf(fields map[string]interface{}, format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	JError(fields map[string]interface{}, v ...interface{})
	JErrorf(fields map[string]interface{}, format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	JFatal(fields map[string]interface{}, v ...interface{})
	JFatalf(fields map[string]interface{}, format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	WithFields(fields map[string]interface{}) Logger
}
