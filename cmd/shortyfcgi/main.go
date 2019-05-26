package main

import (
	"net/http"
	"net/http/fcgi"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/wlbr/shorty/gotils"
	"github.com/wlbr/shorty/net"
)

var config *gotils.Config

func init() {

	cwd, _ := os.Getwd()
	gotils.LogInfo("Current working directory is '%s'.", cwd)

	config = &gotils.Config{}
	gotils.ReadConfig(config)

	gotils.AddAdditionalExpVars(config)
}

func main() {

	gotils.LogDebug("ConfigFile: %s\n", config.ConfigFile)
	gotils.LogDebug("Database: %s\n", config.DataBase.Database)
	gotils.LogDebug("User: %s\n", config.DataBase.User)
	gotils.LogDebug("Password: ***\n") //sorry

	runtime.GOMAXPROCS(runtime.NumCPU()) // use all CPU cores
	n := runtime.NumGoroutine() + 1      // initial number of Goroutines

	// install signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	// Spawn request handler

	go func() {
		err := fcgi.Serve(nil, http.HandlerFunc(net.Handler))
		if err != nil {
			gotils.LogInfo("Not in fcgi mode so not spawing handler.")
			c <- syscall.SIGTERM
			//panic(err)
		} else {
			gotils.LogInfo("Spawing handler.")
		}
	}()

	// give pending requests in fcgi.Serve() some time to enter the request handler
	time.Sleep(time.Millisecond * 100)

	// wait at most 3 seconds for request handlers to finish
	//inc ase we are downloading the GeoIPFile that may take a while
	for i := 0; i < 30; i++ {
		if runtime.NumGoroutine() <= n {
			return
		}
		time.Sleep(time.Millisecond * 100)
	}

	// catch finished handler signal
	_ = <-c
}
