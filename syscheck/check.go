package syscheck

import (
	"privesc-toolkit-manager/types"
	"runtime"
)

func GetSystem() int {
	switch runtime.GOOS {
	case "linux":
		return types.Linux
	case "windows":
		return types.Windows
	case "darwin":
		return types.Mac
	}
	panic("Unknown OS")
}
