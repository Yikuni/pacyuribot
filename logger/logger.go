package logger

import "log"

var (
	DEBUG = true
)

const (
	Reset  = "\033[0m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
)

func Debug(format string, v ...any) {
	if !DEBUG {
		return
	}
	log.Printf("[DEBUG]  : "+format+"\n", v...)
}

func Info(format string, v ...any) {
	log.Printf("[INFO]   : "+format+"\n", v...)
}

func Warning(format string, v ...any) {
	log.Printf(Yellow+"[WARNING]: "+format+Reset+"\n", v...)
}

func Error(format string, v ...any) {
	log.Printf(Red+"[ERROR]  : "+format+Reset+"\n", v...)
}

func Fatal(format string, v ...any) {
	log.Fatalf(Red+"[ERROR]  : "+format+Reset+"\n", v...)
}
