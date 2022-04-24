package fs

import (
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	"bufio"
	"fmt"
	"os"
)

func exists(p string) bool {
	_, e := os.Stat(p)
	return !os.IsNotExist(e)
}

func readFileLines(p string) ([]string, error) {
	f, e := os.Open(p)

	if e != nil {
		return []string{}, e
	}

	var l []string

	s := bufio.NewScanner(f)

	for s.Scan() {
		l = append(l, s.Text())
	}

	return l, s.Err()
}

func Register() {
	require.RegisterNativeModule("fs", func(runtime *engine.Runtime, module *engine.Object) {
		o := module.Get("exports").(*engine.Object)

		o.Set("appendFileLines", func(p string, l []string) error {
			f, e := os.Open(p)

			if e != nil {
				return e
			}

			w := bufio.NewWriter(f)

			for _, cl := range l {
				fmt.Fprintln(w, cl)
			}

			f.Close()

			return w.Flush()
		})

		o.Set("appendFile", func(p string, d string) error {
			f, e := os.Open(p)

			if e != nil {
				return e
			}

			w := bufio.NewWriter(f)

			_, e = w.WriteString(d)

			if e != nil {
				return e
			}

			f.Close()

			return w.Flush()
		})

		o.Set("writeFileLines", func(p string, l []string) error {
			f, e := os.Create(p)

			if e != nil {
				return e
			}

			w := bufio.NewWriter(f)

			for _, cl := range l {
				fmt.Fprintln(w, cl)
			}

			f.Close()

			return w.Flush()
		})

		o.Set("writeFile", func(p string, d string) error {
			return os.WriteFile(p, []byte(d), 0666)
		})

		o.Set("readFileLines", func(p string) (engine.Value, error) {
			l, e := readFileLines(p)
			return runtime.ToValue(l), e
		})

		o.Set("readFile", func(p string) (engine.Value, error) {
			d, e := os.ReadFile(p)
			return runtime.ToValue(string(d)), e
		})

		o.Set("exists", func(p string) (engine.Value) {
			return runtime.ToValue(exists(p))
		})
	})
}
