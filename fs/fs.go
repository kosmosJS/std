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

func writeFile(p string, l []string) error {
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
}

func readFile(p string) ([]string, error) {
	f, e := os.Open(p)

	if e != nil {
		return []string{}, e
	}

	var l []string

	s := bufio.NewScanner(f)

	for s.Scan() {
		l = append(l, s.Text())
	}

	f.Close()

	return l, s.Err()
}

func copyFile(src string, dst string) error {
	d, e := readFile(src)
	if e != nil {
		return d, e
	}
	e = writeFile(dst, d)
	return e
}

func Register() {
	require.RegisterNativeModule("fs", func(runtime *engine.Runtime, module *engine.Object) {
		o := module.Get("exports").(*engine.Object)

		o.Set("appendFile", func(p string, l []string) error {
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

		o.Set("writeFile", func(p string, l []string) error {
			e := writeFile(p, l)
			return e
		})

		o.Set("readFile", func(p string) (engine.Value, error) {
			d, e := readFile(p)
			return runtime.ToValue(d), e
		})

		o.Set("copyFile", func(src string, dst string), error {
			return copyFile(src, dst)
		})

		o.Set("exists", func(p string) (engine.Value) {
			return runtime.ToValue(exists(p))
		})
	})
}
