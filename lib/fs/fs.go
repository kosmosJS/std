package fs

import (
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	cp "github.com/otiai10/copy"
	"errors"
	"bufio"
	"io/fs"
	"fmt"
	"os"
)

func appendFile(p string, l []string) error {
	if isDir(p) {
		return errors.New("'appendFile' cannot be run on a directory.")
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
}

func writeFile(p string, l []string, m os.FileMode) error {
	if isDir(p) {
		return errors.New("'writeFile' cannot be run on a directory.")
	}

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
	if isDir(p) {
		return []string{}, errors.New("'readFile' cannot be run on a directory.")
	}

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

func exists(p string) bool {
	_, e := os.Stat(p)
	return !os.IsNotExist(e)
}

func isDir(p string) bool {
	i, e := os.Stat(p)

	if e != nil {
		return false
	}

	return i.IsDir()
}

func isFile(p string) bool {
	if exists(p) && !isDir(p) {
		i, e := os.Stat(p)

		if e != nil || i.Mode()&fs.ModeSymlink == 0 {
			return false
		}
	}

	return false
}

func isSymlink(p string) bool {
	if exists(p) && !isDir(p) {
		i, e := os.Stat(p)


		if e != nil {
			return false
		} else if i.Mode()&fs.ModeSymlink != 0 {
			return true
		}
	}

	return false
}

func ensureFile(p string, m os.FileMode) error {
	if !exists(p) {
		return writeFile(p, []string{}, m)
	}

	return nil
}

func ensureDir(p string, m os.FileMode) error {
	if !exists(p) {
		return os.MkdirAll(p, m)
	} else if !isDir(p) {
		return errors.New("'" + p + "' already exists, but isn't a directory.")
	}

	return nil
}

func ensureSymlink(s, d string) error {
	if !exists(d) {
		return os.Symlink(s, d)
	}

	i, e := os.Lstat(d)

	if e != nil {
		return e
	}

	if i.Mode()&fs.ModeSymlink == 0 {
		return errors.New("'" + d + "' already exists, but isn't a symlink.")
	}

	return nil
}

func Register() {
	require.RegisterNativeModule("fs", func(runtime *engine.Runtime, module *engine.Object) {
		o := module.Get("exports").(*engine.Object)

		o.Set("appendFile", func(p string, l []string) error {
			return appendFile(p, l)
		})

		o.Set("writeFile", func(p string, l []string, m os.FileMode) error {
			return writeFile(p, l, m)
		})

		o.Set("readFile", func(p string) (engine.Value, error) {
			d, e := readFile(p)
			return runtime.ToValue(d), e
		})

		o.Set("rename", func(s, d string) error {
			return os.Rename(s, d)
		})

		o.Set("copy", func(s, d string) error {
			return cp.Copy(s, d)
		})

		o.Set("remove", func(p string) error {
			if isDir(p) {
				return os.RemoveAll(p)
			}

			return os.Remove(p)
		})

		o.Set("isFile", func(p string) engine.Value {
			return runtime.ToValue(isDir(p))
		})

		o.Set("isDir", func(p string) engine.Value {
			return runtime.ToValue(isDir(p))
		})

		o.Set("isFile", func(p string) engine.Value {
			return runtime.ToValue(isFile(p))
		})

		o.Set("isSymlink", func(p string) engine.Value {
			return runtime.ToValue(isSymlink(p))
		})

		o.Set("mkdir", func(p string, m os.FileMode) error {
			return os.MkdirAll(p, m)
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

		o.Set("createSymlink", func(s, d string) error {
			return os.Symlink(s, d)
		})

		o.Set("ensureFile", func(p string, m os.FileMode) error {
			return ensureFile(p, m)
		})

		o.Set("ensureDir", func(p string, m os.FileMode) error {
			return ensureDir(p, m)
		})

		o.Set("ensureSymlink", func(s, d string) error {
			return ensureSymlink(s, d)
		})
	})
}
