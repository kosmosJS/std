package process

import (
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	"strings"
	"os"
)

func env() map[string]string {
	var v map[string]string

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		v[pair[0]] = pair[1]
	}

	return v
}

func Register() {
	require.RegisterNativeModule("process", func(runtime *engine.Runtime, module *engine.Object) {
		o := module.Get("exports").(*engine.Object)

		o.Set("exit", func(c int) {
			os.Exit(c)
		})

		o.Set("args", func() engine.Value {
			return runtime.ToValue(os.Args[1:])
		})

		o.Set("env", func() engine.Value {
			return runtime.ToValue(env())
		})

		o.Set("chdir", func(p string) error {
			return os.Chdir(p)
		})

		o.Set("cwd", func() (engine.Value, error) {
			d, e := os.Getwd()
			return runtime.ToValue(d), e
		})
	})
}