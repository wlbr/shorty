package gotils

import (
	"expvar"
	"os"
	"runtime"
	"strings"
	"time"
)

// AddAdditionalExpVars extends the standard expvars (see https://golang.org/pkg/expvar/ )
// by some environment variables and config data.
func AddAdditionalExpVars(config *Config) {
	// static values
	expvar.Publish("GOROOT", expvar.Func(func() interface{} { return runtime.GOROOT() }))
	expvar.Publish("GOOS", expvar.Func(func() interface{} { return runtime.GOOS }))
	expvar.Publish("GOARCH", expvar.Func(func() interface{} { return runtime.GOARCH }))
	expvar.Publish("Compiler", expvar.Func(func() interface{} { return runtime.Compiler }))
	expvar.Publish("CompilerVersion", expvar.Func(func() interface{} { return runtime.Version() }))
	expvar.Publish("BuildVersion", expvar.Func(func() interface{} { return config.GitVersion }))
	expvar.Publish("BuildTimeStamp", expvar.Func(func() interface{} { return config.BuildTimeStamp.String() }))
	expvar.Publish("FullComandline", expvar.Func(func() interface{} { return strings.Join(os.Args, " ") }))
	/*expvar.Publish("BuildInfo", expvar.Func(func() interface{} {
		bi, ok := debug.ReadBuildInfo()
		if ok {
			return fmt.Sprintf("path: %s, deps: %v main: %v", bi.Path, bi.Deps, bi.Main)
		}
		return ""
	}))*/

	//dynamic values
	expvar.Publish("NumCPU", expvar.Func(func() interface{} { return runtime.NumCPU() }))
	expvar.Publish("NumGoroutine", expvar.Func(func() interface{} { return runtime.NumGoroutine() }))
	expvar.Publish("ActiveLogLevel", expvar.Func(func() interface{} { return config.Log.ActiveLogLevel.String() }))
	expvar.Publish("LogFile", expvar.Func(func() interface{} { return config.Log.LogFileName }))
	expvar.Publish("CurrentTime", expvar.Func(func() interface{} { return time.Now() }))
}
