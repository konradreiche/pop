package pop

import "github.com/sirupsen/logrus"

// AvailableDialects lists the available database dialects
var AvailableDialects = []string{"postgres", "mysql", "cockroach"}

// Debug mode, to toggle verbose log traces
var Debug = false

// Color mode, to toggle colored logs
// Deprecated: Use pop.Logger instead.
var Color = true

// Log a formatted string to the logger
// Deprecated: Use pop.Logger.Debugf() or pop.Logger.Debug() instead.
var Log = func(s string, args ...interface{}) {
	if Debug {
		Logger.Infof(s, args)
	}
}

var Logger = logrus.WithField("component", "pop")

// DialectSupported checks support for the given database dialect
func DialectSupported(d string) bool {
	for _, ad := range AvailableDialects {
		if ad == d {
			return true
		}
	}
	return false
}
