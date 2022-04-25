package os

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	"github.com/pbnjay/memory"
	rt "runtime"
	"os"
)

func isPosix() {
	if rt.GOOS != "windows" {
		return true
	}

	return false
}

func getEndianness() string {
	var i int = 0x1
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		return "big"
	} else {
		return "little"
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

		o.Set("eol", runtime.ToValue(isPosix() ? "\n" : "\r\n"))

		o.Set("arch", runtime.ToValue(rt.GOARCH))

		o.Set("platform", runtime.ToValue(rt.GOOS))

		o.Set("null", runtime.ToValue(isPosix() ? "/dev/null" : "\\.\nul"))

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
