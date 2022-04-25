package fs

import (
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	cp "github.com/otiai10/copy"
	"errors"
	"bufio"
	"fmt"
	"os"
)

func exists(p string) bool {
	_, e := os.Stat(p)
	return !os.IsNotExist(e)
}

func writeFile(p string, l []string, m os.FileMode) error {
	f, e := os.Create(p)

	if e != nil {
		return e
	}

	e = f.Chmod(m)

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

func isDir(p string) bool {
	i, e := os.Stat(p)
	if e != nil {
		return false
	}
	return i.IsDir()
}

func Register() {
	require.RegisterNativeModule("fs", func(runtime *engine.Runtime, module *engine.Object) {
		o := module.Get("exports").(*engine.Object)

		o.Set("append", func(p string, l []string) error {
			if isDir(p) {
				return errors.New("'append' cannot be run on a directory.")
			}

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

		o.Set("write", func(p string, l []string, m os.FileMode) error {
			if isDir(p) {
				return errors.New("'write' cannot be run on a directory.")
			}

			e := writeFile(p, l, m)
			return e
		})

		o.Set("read", func(p string) (engine.Value, error) {
			if isDir(p) {
				return runtime.ToValue([]string{}), errors.New("'read' cannot be run on a directory.")
			}

			d, e := readFile(p)
			return runtime.ToValue(d), e
		})

		o.Set("rename", func(src string, dst string) error {
			return os.Rename(src, dst)
		})

		o.Set("copy", func(src string, dst string) error {
			return cp.Copy(src, dst)
		})

		o.Set("remove", func(p string) error {
			if isDir(p) {
				return os.RemoveAll(p)
			}

			return os.Remove(p)
		})

		o.Set("isDir", func(p string) engine.Value {
			return runtime.ToValue(isDir(p))
		})

		o.Set("mkdir", func(p string, m os.FileMode) error {
			return os.MkdirAll(p, m)
		})

		o.Set("chdir", func(p string) error {
			return os.Chdir(p)
		})

		o.Set("chmod", func(p string, m os.FileMode) error {
			return os.Chmod(p, m)
		})

		o.Set("chown", func(p string, u, g int) error {
			return os.Chown(p, u, g)
		})

		o.Set("exists", func(p string) engine.Value {
			return runtime.ToValue(exists(p))
		})
	})
}
