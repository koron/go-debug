// +build !debug

package debug

import "log"

// Logger is output for debug log.
var Logger *log.Logger

// Print outputs debug log if built with "debug" tag. Arguments are handled in
// the manner of fmt.Print.
func Print(v ...interface{}) {
	// nothing to output for release.
}

// Printf outputs debug log if built with "debug" tag. Arguments are handled in
// the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	// nothing to output for release.
}

// Println calls debug log if built with "debug" tag. Arguments are handled in
// the manner of fmt.Println.
func Println(v ...interface{}) {
	// nothing to output for release.
}
