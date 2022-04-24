package fs

import (
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	//"os"
)

func RegisterFS() {
	require.RegisterNativeModule("fs", func(runtime *engine.Runtime, module *engine.Object) {
		o := module.Get("exports").(*engine.Object)

		o.Set("readFile", func(call engine.FunctionCall, name string) engine.Value {
			return runtime.ToValue(name)
		})
	})
}
