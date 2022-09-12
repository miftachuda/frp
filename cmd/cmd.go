package cmd

import (
	"runtime"

	proxy "github.com/bdwyertech/gontlm-proxy/pkg"
)

func Execute() {
	if runtime.GOOS == "windows" {
		proxy.RunWindows()
	} else {
		proxy.Run()
	}
}
