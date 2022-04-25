package os

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	"github.com/pbnjay/memory"
	rt "runtime"
	"unsafe"
	"os"
)

func isPosix() bool {
	if rt.GOOS != "windows" {
		return true
	}

	return false
}

func getEndianness() string {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch (buf) {
	case [2]byte{0xCD, 0xAB}:
		return "little"
	case [2]byte{0xAB, 0xCD}:
		return "big"
	default:
		return "unknown"
	}
}

func Register() {
	require.RegisterNativeModule("os", func(runtime *engine.Runtime, module *engine.Object) {
		osType := ""

		switch (rt.GOOS) {
		case "aix":
			osType = "AIX"
		case "android":
			osType = "Linux"
		case "dragonfly":
			osType = "DragonFly"
		case "darwin":
			osType = "Darwin"
		case "illumos":
			osType = "SunOS"
		case "ios":
			osType = "Darwin"
		case "linux":
			osType = "Linux"
		case "netbsd":
			osType = "NetBSD"
		case "openbsd":
			osType = "OpenBSD"
		case "solaris":
			osType = "SunOS"
		case "windows":
			osType = "Windows_NT"
		}

		o := module.Get("exports").(*engine.Object)

		var eol string
		if isPosix() {
			eol = "\n"
		} else {
			eol = "\r\n"
		}

		var null string
		if isPosix() {
			null = "/dev/null"
		} else {
			null = "\\.\nul"
		}

		o.Set("eol", runtime.ToValue(eol))

		o.Set("arch", runtime.ToValue(rt.GOARCH))

		o.Set("platform", runtime.ToValue(rt.GOOS))

		o.Set("null", runtime.ToValue(null))

		o.Set("endianness", runtime.ToValue(getEndianness()))

		o.Set("posix", runtime.ToValue(isPosix()))

		o.Set("type", runtime.ToValue(osType))

		o.Set("memtotal", func() engine.Value {
			return runtime.ToValue(memory.TotalMemory())
		})

		o.Set("memfree", func() engine.Value {
			return runtime.ToValue(memory.FreeMemory())
		})

		o.Set("memused", func() engine.Value {
			return runtime.ToValue(memory.TotalMemory() - memory.FreeMemory())
		})

		o.Set("hostname", func() (engine.Value, error) {
			n, e := os.Hostname()
			return runtime.ToValue(n), e
		})
		
		o.Set("homedir", func() (engine.Value, error) {
			d, e := homedir.Dir()
			return runtime.ToValue(d), e
		})
	})
}
