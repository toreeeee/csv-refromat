package console

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func Error(format string, a ...any) {
	fmt.Printf("%sError: %s%s\n", Red, fmt.Sprintf(format, a...), Reset)
}

func Success(format string, a ...any) {
	fmt.Printf("%sSuccess: %s%s\n", Green, fmt.Sprintf(format, a...), Reset)
}

func Info(format string, a ...any) {
	fmt.Printf("%sInfo: %s%s\n", Blue, fmt.Sprintf(format, a...), Reset)
}

func Warn(format string, a ...any) {
	fmt.Printf("%sWarning: %s%s\n", Yellow, fmt.Sprintf(format, a...), Reset)
}
