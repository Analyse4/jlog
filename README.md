# Overview
jlog is a simple encapsulation for adding log level function against the standard library log package.
# Installing
Using jlog is easy. First, use go get to install the latest version of the library:
```bash
go get -u github.com/Analyse4/jlog
```
Next, include jlog in your application:
```go
import "github.com/Analyse4/jlog"
```
# Getting Started
You need init in main package firstly when you about to use jlog as your project log package.

For example main.go
```go
func init() {
	jlog.Init(os.Stdout, "", jlog.LstdFlags|jlog.Lshortfile)
}

func main() {
	// main func code
}
```
Then you need set current log level.

`jlog.LstdFlags|jlog.Lshortfile` is used to make jlog print data, time, file and line number:
```go
func init() {
    jlog.Init(os.Stdout, "", jlog.LstdFlags|jlog.Lshortfile)
    jlog.SetLevel(jlog.DEBUG)
}

func main() {
	// main func code
}
```
Finally you can use jlog for logging anywhere.

For example:
```go
jlog.Info("hello, jlog")
jlog.Debugf("%s\n", "hello jog")
```
Output:
```bash
[INFO] 2020-06-15T14:36:20+08:00 main.go:18: hello jlog
[DEBUG] 2020-06-15T14:36:20+08:00 main.go:19: hello jlog
```
Specifically detail in [doc]()