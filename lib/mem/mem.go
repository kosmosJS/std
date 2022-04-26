package mem

import (
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	"github.com/pbnjay/memory"
	rt "runtime"
)

func process() map[string]map[string]int {
	var m rt.MemStats
	rt.ReadMemStats(&m)

	ms := map[string]map[string]int {
		"gc": {
			"cycles": int(m.NumGC),
			"forcedCycles": int(m.NumForcedGC),
			"nextSize": int(m.NextGC),
			"lastTime": int(m.LastGC),
		},
		"heap": {
			"alloc": int(m.Alloc),
			"totalAlloc": int(m.TotalAlloc),
			"objects": int(m.HeapObjects),
			"obtained": int(m.HeapSys),
			"released": int(m.HeapReleased),
			"active": int(m.HeapInuse),
			"inactive": int(m.HeapIdle),
		},
	}

	return ms
}

func Register() {
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

		o.Set("forceGC", func() {
			rt.GC()
		})
	})
}