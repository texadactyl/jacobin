//go:build windows

package native

import (
	"fmt"
	"golang.org/x/sys/windows"
	"jacobin/trace"
)

func ConnectLibrary(libPath string) uintptr {
	handle, err := windows.LoadLibrary(libPath)
	if err != nil {
		errMsg := fmt.Sprintf("ConnectLibrary: windows.LoadLibrary for [%s] failed, reason: [%s]",
			libPath, err.Error())
		trace.Error(errMsg)
		handle = 0
	}
	return uintptr(handle)
}
