package utils

import (
	"runtime/debug"
	"strings"
)

var (
	appHash string
)

func init() {
	GetLastCommitHash()
}

func GetLastCommitHash() string {
	if appHash != "" {
		return appHash
	}
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	data := strings.Split(strings.ReplaceAll(info.Main.Version, "+dirty", ""), "-")
	res := data[len(data)-1]
	if len(res) > 7 {
		return res[:7]
	}
	appHash = res
	return appHash
}
