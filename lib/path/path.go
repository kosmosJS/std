package path

import (
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	"path/filepath"
	"strings"
	"regexp"
	"errors"
)

func dir(p string, sep string) string {
	pp, sp := split(normalize(p, sep), sep)

	return pp + strings.Join(sp[:len(sp)-1], sep)
}

func format(d map[string]string, sep string) (string, error) {
	var dir string
	var name string
	var ext string
	var ok bool

	dir, ok = d["dir"]

	if !ok {
		return "", errors.New("dir is not defined")
	}

	name, ok = d["name"]

	if !ok {
		return "", errors.New("name is not defined")
	}

	ext, ok = d["ext"]

	if !ok {
		return "", errors.New("ext is not defined")
	}

	p := normalize(dir + sep + name + ext, sep)

	return p, nil
}

func parse(p string, sep string) map[string]string {
	data := map[string]string {
		"dir": dir(p, sep),
		"name": name(p, sep),
		"ext": ext(p, sep),
	}

	return data
}

func split(p string, sep string) (string, []string) {
	r, _ := regexp.Compile("^.:\\")

	var pp string

	for strings.HasPrefix(p, "file://") {
		pp += "file://"
		p = strings.Replace(p, "file://", "", 1)
	}

	for strings.HasPrefix(p, sep) {
		pp += sep
		p = strings.Replace(p, sep, "", 1)
	}

	for r.MatchString(p) {
		pp += r.FindString(p)
		p = r.ReplaceAllString(sep, "")
	}

	for strings.HasSuffix(p, sep) {
		p = strings.Replace(p, sep, "", -1)
	}

	sp := strings.Split(p, sep)

	return pp, sp
}

func normalize(p string, sep string) string {
	pp, sp := split(p, sep)

	var ns []string

	for _, i := range sp {
		if i != "" && i != ".." {
			ns = append(ns, i)
			continue
		}

		ns = ns[:len(ns) - 1]
	}

	fp, _ := filepath.Abs(pp + strings.Join(ns, sep))

	return fp
}

func isAbsolute(p string, sep string) bool {
	p = normalize(p)

	r, _ := regexp.Compile("^.:\\")
	
	if strings.HasPrefix(p, sep) || strings.HasPrefix(p, "file://"+sep) || r.MatchString(p) {
		return true
	}

	return false
}

func fromFileURL(p string, sep string) string {
	for strings.HasPrefix(p, "file://") {
		p = strings.Replace(p, "file://", "", 1)
	}

	return normalize(p)
}

func toFileURL(p string, sep string) string {
	p = normalize(p)

	for !strings.HasPrefix(p, "file://") {
		p = "file://" + p
	}

	return p
}

func join(l []string, sep string) string {
	return normalize(strings.Join(l, sep), sep)
}

func base(p string, sep string) string {
	return strings.Replace(p, dir(p, sep), "", 1)
}

func name(p string, sep string) string {
	return strings.Replace(strings.Replace(p, dir(p, sep), "", 1), ext(p, sep), "", 1)
}

func ext(p string, sep string) string {
	sp := strings.Split(strings.Replace(p, dir(p, sep), "", 1), ".")

	return strings.Join(sp[1:len(sp)-1], ".")
}

func Register(px bool) {
	require.RegisterNativeModule("os", func(runtime *engine.Runtime, module *engine.Object) {
		o := module.Get("exports").(*engine.Object)

		var sep string
		if px {
			sep = "/"
		} else {
			sep = `\`
		}

		var d string
		if px {
			d = ":"
		} else {
			d = ";"
		}

		o.Set("sep", runtime.ToValue(sep))

		o.Set("delimiter", runtime.ToValue(d))

		o.Set("dir", func(p string) engine.Value {
			return runtime.ToValue(dir(p, sep))
		})

		o.Set("ext", func(p string) engine.Value {
			return runtime.ToValue(ext(p, sep))
		})

		o.Set("format", func(d map[string]string) (engine.Value, error) {
			f, e := format(d, sep)
			return runtime.ToValue(f), e
		})

		o.Set("parse", func(p string) engine.Value {
			return runtime.ToValue(parse(p, sep))
		})

		o.Set("isAbsolute", func(p string) engine.Value {
			return runtime.ToValue(isAbsolute(p, sep))
		})

		o.Set("base", func(p string) engine.Value {
			return runtime.ToValue(base(p, sep))
		})

		o.Set("name", func(p string) engine.Value {
			return runtime.ToValue(name(p, sep))
		})

		o.Set("join", func(l []string) engine.Value {
			return runtime.ToValue(join(l, sep))
		})

		o.Set("normalize", func(p string) engine.Value {
			return runtime.ToValue(normalize(p, sep))
		})

		o.Set("split", func(p string) engine.Value {
			pp, sp := split(p, sep)
			return runtime.ToValue([]interface{}{pp, sp})
		})
	})
}
