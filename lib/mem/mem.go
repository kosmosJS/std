package mem

import (
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	"github.com/pbnjay/memory"
	rt "runtime"
	"debug"
	"os"
)

func process() map[string]map[string]string {
	var m rt.MemStats
	rt.ReadMemStats(&m)

	ms := map[string]map[string]string {
		"gc": {
			"cycles": m.NumGC,
			"forcedCycles": m.NumForcedGC,
			"nextSize": m.NextGC,
			"lastTime": m.LastGC,
		},
		"heap": {
			"alloc": m.Alloc,
			"totalAlloc": m.TotalAlloc,
			"objects": m.HeapObjects,
			"obtained": m.HeapSys,
			"released": m.HeapReleased,
			"active": m.HeapInuse,
			"inactive": m.HeapIdle,
		},
	}

	return ms
}

func Register(px bool) {
	var gcp int

	if os.Getenv("GOGC") {
		gcp = os.Getenv("GOGC")
	} else {
		gcp = 100
	}

	require.RegisterNativeModule("mem", func(runtime *engine.Runtime, module *engine.Object) {
		o := module.Get("exports").(*engine.Object)

		o.Set("total", func() engine.Value {
			return runtime.ToValue(memory.TotalMemory())
		})

		o.Set("free", func() engine.Value {
			return runtime.ToValue(memory.FreeMemory())
		})

		o.Set("used", func() engine.Value {
			return runtime.ToValue(memory.TotalMemory() - memory.FreeMemory())
		})

		o.Set("process", func() engine.Value {
			return runtime.ToValue(process())
		})

		o.Set("forceAggressiveGC", func() {
			debug.FreeOSMemory()
		})

		o.Set("forceGC", func() {
			rt.GC()
		})

		o.Set("setGCPercent", func(p int) {
			gcp = p
			debug.setGCPercent(p)
		})

		o.Set("getGCPercent", func(p int) engine.Value {
			return runtime.ToValue(gcp)
		})
	})
}