package ghost

import (
	"strconv"
	"sync/atomic"

	"github.com/yimikao/ghost/core/severity"
)

type logger struct {
	stderrThreshold severityValue
}

var globalLogger logger

// severity identifies the sort of log: info, warning etc. It also implements
// flag.Value interface. The -stderrthreshold flag is of type severity and
// should be modified only through the flag.Value interface.
type severityValue struct {
	severity.Severity
}

// get returns the value of wrapped Severity.
func (s *severityValue) get() severity.Severity {
	return severity.Severity(atomic.LoadInt32((*int32)(&s.Severity)))
}

//set sets the value of wrapped Severity.
func (s *severityValue) set(value severity.Severity) {
	atomic.StoreInt32((*int32)(&s.Severity), int32(value))
}

// String is part of the flag.Value interface.
func (s *severityValue) String() string {
	return strconv.FormatInt(int64(s.Severity), 10)
}

// Get is part of the flag.Getter interface
func (s *severityValue) Get() severity.Severity {
	return s.Severity
}

// Set is part of the flag.Value interface.
// set severity threshold by value passed.
func (s *severityValue) Set(value string) error {
	var threshold severity.Severity
	//check if passed value matches any defined severity.
	// CASE "INFO" etc
	if sev, ok := severity.ByString(value); ok {
		threshold = sev
	} else {
		// CASE "10" etc
		sev, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		threshold = severity.Severity(sev)
	}
	globalLogger.stderrThreshold.set(threshold)
	return nil
}
