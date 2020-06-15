package jlog

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Log level.
const (
	INFO  = 200
	DEBUG = 201
)

// Log flag
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

// jl is singleton pattern log instance for other package to use.
var jl *JLogger

// A JLogger represents an active logging object that generates lines of
// output to an io.Writer. Each logging operation makes a single call to
// the Writer's Write method. A JLogger can be used simultaneously from
// multiple goroutines; it guarantees to serialize access to the Writer.
type JLogger struct {
	stdlog *log.Logger
	level  int
	flag   int
}

// Init creates a new JLogger. The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line, or
// after the log header if the Lmsgprefix flag is provided.
// The flag argument defines the logging properties.
func Init(out io.Writer, prefix string, flag int) {
	jl = new(JLogger)
	jl.stdlog = log.New(out, prefix, 0)
	jl.flag = flag
}

// // Get return jlogger for log.
// func Get() (*JLogger, error) {
// 	if jl == nil {
// 		return nil, fmt.Errorf("jlogger is nil")
// 	}
// 	return jl, nil
// }

// SetLevel set log level.
func SetLevel(level int) {
	jl.level = level
}

// Info only print logs prefixed by INFO.
// Info calls jl.stdlog.Println to print the logger.
// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func Info(v ...interface{}) {
	var s string
	jl.stdlog.SetPrefix("[INFO] ")
	if jl.flag == LstdFlags|Lshortfile {
		s = generateStdflagShortFile()
	}

	s = s + fmt.Sprintln(v...)
	jl.stdlog.Print(s)
}

// Debug only print logs prefixed by DEBUG and higher.
// Debug calls jl.stdlog.Println to print the logger.
// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func Debug(v ...interface{}) {
	if jl.level != INFO {
		var s string
		jl.stdlog.SetPrefix("[DEBUG] ")
		if jl.flag == LstdFlags|Lshortfile {
			s = generateStdflagShortFile()
		}

		s = s + fmt.Sprintln(v...)
		jl.stdlog.Print(s)
	}
}

// Infof calls stdlog.Printf to print to the logger.
// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	var s string
	jl.stdlog.SetPrefix("[INFO] ")
	if jl.flag == LstdFlags|Lshortfile {
		s = generateStdflagShortFile()
	}

	jl.stdlog.Printf(s+format, v...)
}

// Debugf calls stdlog.Printf to print to the logger.
// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	if jl.level != INFO {
		var s string
		jl.stdlog.SetPrefix("[DEBUG] ")
		if jl.flag == LstdFlags|Lshortfile {
			s = generateStdflagShortFile()
		}

		jl.stdlog.Printf(s+format, v...)
	}
}

// generateStdflagShortFile generate data, time, file and line number prefix.
func generateStdflagShortFile() string {
	var s string
	var sf string

	_, fn, ln, ok := runtime.Caller(2)
	if ok {
		sf = fn + ":" + strconv.Itoa(ln)
		index := strings.LastIndex(sf, "/")
		sf = sf[index+1:]
	}
	s = time.Now().Format(time.RFC3339) + " " + sf + ": "
	return s
}
