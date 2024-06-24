package console

import "fmt"

const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	gray    = "\033[37m"
	white   = "\033[97m"
)

func Error(format string, a ...any) {
	fmt.Printf("%sError: %s%s\n", red, fmt.Sprintf(format, a...), reset)
}

func Success(format string, a ...any) {
	fmt.Printf("%sSuccess: %s%s\n", green, fmt.Sprintf(format, a...), reset)
}

func Info(format string, a ...any) {
	fmt.Printf("%sInfo: %s%s\n", blue, fmt.Sprintf(format, a...), reset)
}

func Warn(format string, a ...any) {
	fmt.Printf("%sWarning: %s%s\n", yellow, fmt.Sprintf(format, a...), reset)
}

func Debug(format string, a ...any) {
	fmt.Printf("%sWarning: %s%s\n", gray, fmt.Sprintf(format, a...), reset)
}
