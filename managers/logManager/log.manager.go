package logManager

import (
	"github.com/astaxie/beego/logs"
)

// LogInfo logs Info messages for the code
func LogInfo(f interface{}, objects ...interface{}) {
	logs.Info(f, objects...)
}

// LogError logs Error messages for the code
func LogError(f interface{}, objects ...interface{}) {
	logs.Error(f, objects...)
}

// LogDebug logs debug messages for the code
func LogDebug(f interface{}, objects ...interface{}) {
	logs.Debug(f, objects...)
}
