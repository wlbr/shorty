package gotils

import (
	"fmt"
	"runtime"
)

//CheckErr is a convenience function makes error handling dangerously simple.
func CheckErr(err error) {
	if err != nil {
		LogError("%s", err)
	}
}

// GetInspectData offer some additional debugging information
func GetInspectData(config Config) string {
	return fmt.Sprintf("Version: %s \n", config.GitVersion) +
		fmt.Sprintf("Build timestamp: %s \n", config.BuildTimeStamp) +
		fmt.Sprintf("Golang compile version: %s \n", runtime.Version()) +
		fmt.Sprintf("Compile GOROOT: %s \n", runtime.GOROOT()) +
		fmt.Sprintf("Compile GOOS: %s \n", runtime.GOOS) +
		fmt.Sprintf("Compile GOARCH: %s \n", runtime.GOARCH) +
		fmt.Sprintf("Runtime NumCPU: %d \n", runtime.NumCPU()) +
		fmt.Sprintf("Runtime NumGoroutine: %d \n", runtime.NumGoroutine()) +
		fmt.Sprintf("ActiveLogLevel: %d \n", config.Log.ActiveLogLevel) +
		fmt.Sprintf("LogFile: %s \n", config.Log.LogFileName)
}

// Minf64 returns the minimum of a slice of float64
func Minf64(v []float64) float64 {
	m := v[0]
	for _, e := range v {
		if e < m {
			m = e
		}
	}
	return m
}

// Maxf64 returns the maximum of a slice of float64
func Maxf64(v []float64) float64 {
	m := v[0]
	for _, e := range v {
		if e > m {
			m = e
		}
	}
	return m
}
