// +build debug

package debug

import (
	"log"
	"os"
)

// Logger is output for debug log.
var Logger = log.New(os.Stdout, "DEBUG ", log.LstdFlags)

// Print outputs debug log if built with "debug" tag. Arguments are handled in
// the manner of fmt.Print.
func Print(v ...interface{}) {
	Logger.Print(v...)
}

// Printf outputs debug log if built with "debug" tag. Arguments are handled in
// the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}

// Println calls debug log if built with "debug" tag. Arguments are handled in
// the manner of fmt.Println.
func Println(v ...interface{}) {
	Logger.Println(v...)
}
