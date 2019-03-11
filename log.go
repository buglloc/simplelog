package log

var (
	logger = NewLogger()
)

func SetLevel(level Lvl) {
	maxLvl = level
}

func Debug(msg string, ctx ...interface{}) {
	logger.Debug(msg, ctx...)
}

func Info(msg string, ctx ...interface{}) {
	logger.Info(msg, ctx...)
}

func Warn(msg string, ctx ...interface{}) {
	logger.Warn(msg, ctx...)
}

func Error(msg string, ctx ...interface{}) {
	logger.Error(msg, ctx...)
}

func Crit(msg string, ctx ...interface{}) {
	logger.Crit(msg, ctx...)
}

func Child(ctx ...interface{}) Logger {
	return logger.Child(ctx...)
}
