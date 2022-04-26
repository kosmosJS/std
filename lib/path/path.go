package path

import (
	"github.com/kosmosJS/engine-node/require"
	"github.com/kosmosJS/engine"
	"path/filepath"
	"strings"
	"regexp"
	"errors"
	"mime"
)

func dir(p, sep string) string {
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

	return join([]string{dir, name + ext}, sep), nil
}

func parse(p, sep string) map[string]string {
	data := map[string]string {
		"dir": dir(p, sep),
		"name": name(p, sep),
		"ext": ext(p, sep),
	}

	return data
}

func split(p, sep string) (string, []string) {
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

func normalize(p, sep string) string {
	pp, sp := split(p, sep)

	var ns []string

	for _, i := range sp {
		if i != "" && i != ".." && i != "." {
			ns = append(ns, i)
			continue
		}

		if i == "." {
			continue
		}

		ns = ns[:len(ns) - 1]
	}

	fp, _ := filepath.Abs(pp + strings.Join(ns, sep))

	return fp
}

func isAbsolute(p, sep string) bool {
	p = normalize(p, sep)

	r, _ := regexp.Compile("^.:\\")

	if strings.HasPrefix(p, sep) || strings.HasPrefix(p, "file://"+sep) || r.MatchString(p) {
		return true
	}

	return false
}

func fromFileURL(u, sep string) string {
	for strings.HasPrefix(u, "file://") {
		u = strings.Replace(u, "file://", "", 1)
	}

	return normalize(u, sep)
}

func toFileURL(p, sep string) string {
	p = normalize(p, sep)

	for !strings.HasPrefix(p, "file://") {
		p = "file://" + p
	}

	return p
}

func common(p1, p2, sep string) string {
	p1 = normalize(p1, sep)
	p2 = normalize(p2, sep)
	pp1, sp1 := split(p1, sep)
	pp2, sp2 := split(p2, sep)

	if pp1 != pp2 {
		return ""
	}

	var fp []string

	for i, p := range sp1 {
		if sp2[i] == p {
			fp = append(fp, p)
		} else {
			break
		}
	}

	return join(fp, sep)
}

func join(l []string, sep string) string {
	return normalize(strings.Join(l, sep), sep)
}

func base(p, sep string) string {
	return strings.Replace(p, dir(p, sep), "", 1)
}

func name(p, sep string) string {
	return strings.Replace(strings.Replace(p, dir(p, sep), "", 1), ext(p, sep), "", 1)
}

func ext(p, sep string) string {
	sp := strings.Split(strings.Replace(p, dir(p, sep), "", 1), ".")

	return strings.Join(sp[1:len(sp)-1], ".")
}

func resolve(p1, p2, sep string) string {
	p1 = normalize(p1, sep)
	p2 = normalize(p2, sep)
	p, _ := split(p1, sep)

	c := common(p1, p2, sep)

	if c == "" || c == p {
		return p2
	}

	return normalize(join([]string{p1, p2}, sep), sep)
}

func relative(p1, p2, sep string) string {
	p1 = normalize(p1, sep)
	p2 = normalize(p2, sep)
	p, _ := split(p1, sep)

	c := common(p1, p2, sep)

	if isAbsolute(p2) || c == p {
		return p2
	}

	return normalize(join([]string{"..", strings.Replace(p2, c, "", 1)}, sep), sep)
}

func getMime(p, sep string) string {
	return mime.TypeByExtension(base(normalize(p, sep)))
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

		o.Set("common", func(p1, p2 string) engine.Value {
			return runtime.ToValue(common(p1, p2, sep))
		})

		o.Set("resolve", func(p1, p2 string) engine.Value {
			return runtime.ToValue(resolve(p1, p2, sep))
		})

		o.Set("relative", func(p1, p2 string) engine.Value {
			return runtime.ToValue(relative(p1, p2, sep))
		})

		o.Set("split", func(p string) engine.Value {
			pp, sp := split(p, sep)
			return runtime.ToValue([]interface{}{pp, sp})
		})

		o.Set("fromFileURL", func(u string) engine.Value {
			return runtime.ToValue(fromFileURL(u, sep))
		})

		o.Set("toFileURL", func(p string) engine.Value {
			return runtime.ToValue(toFileURL(p, sep))
		})

		o.Set("getMime", func(p string) engine.Value {
			return runtime.ToValue(getMime(p, sep))
		})
	})
}
