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

func GetArch() int {
	switch runtime.GOARCH {
	case "amd64":
		return types.Amd64
	case "arm64":
		return types.Arm64
	case "386":
		return types.I386
	}
	panic("Unknown Arch")
}
