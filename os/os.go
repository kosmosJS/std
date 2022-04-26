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

func Register(px bool) {
	require.RegisterNativeModule("os", func(runtime *engine.Runtime, module *engine.Object) {
		ot := ""

		switch (rt.GOOS) {
		case "aix":
			ot = "AIX"
		case "android":
			ot = "Linux"
		case "dragonfly":
			ot = "DragonFly"
		case "darwin":
			ot = "Darwin"
		case "illumos":
			ot = "SunOS"
		case "ios":
			ot = "Darwin"
		case "linux":
			ot = "Linux"
		case "netbsd":
			ot = "NetBSD"
		case "openbsd":
			ot = "OpenBSD"
		case "solaris":
			ot = "SunOS"
		case "windows":
			ot = "Windows_NT"
		}

		var eol string
		if px {
			eol = "\n"
		} else {
			eol = "\r\n"
		}

		var n string
		if px {
			n = "/dev/n"
		} else {
			n = "\\.\nul"
		}

		o := module.Get("exports").(*engine.Object)

		o.Set("eol", runtime.ToValue(eol))

		o.Set("arch", runtime.ToValue(rt.GOARCH))

		o.Set("platform", runtime.ToValue(rt.GOOS))

		o.Set("null", runtime.ToValue(n))

		o.Set("endianness", runtime.ToValue(getEndianness()))

		o.Set("posix", runtime.ToValue(px))

		o.Set("type", runtime.ToValue(ot))

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

		o.Set("chdir", func(p string) error {
			return os.Chdir(p)
		})
	})
}
