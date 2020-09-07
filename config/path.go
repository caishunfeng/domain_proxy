package config

import (
	"path"
	"runtime"
)

var (
	ConfigRoot string
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	ConfigRoot = path.Dir(filename)
}
