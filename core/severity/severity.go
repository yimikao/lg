package severity

import "strings"

// Severity identifies the type of log: info, error etc.
type Severity int32

// These constants identify the log levels in order of increasing severity.
const (
	Info Severity = iota
	Warn
	Error
	Fatal
)

// Strings holds all severity levels in string forms indexed by there value
var Strings = []string{
	Info:  "INFO",
	Warn:  "WARN",
	Error: "ERROR",
	Fatal: "FATAL",
}

// ByString returns a severity level by it's string form
func ByString(s string) (Severity, bool) {
	for sev, str := range Strings {
		if str == strings.ToUpper(s) {
			return Severity(sev), true
		}
	}
	return -1, false
}
