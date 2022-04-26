package std

import (
	kPATH "github.com/kosmosJS/std/lib/path"
	kFS "github.com/kosmosJS/std/lib/fs"
	kOS "github.com/kosmosJS/std/lib/os"
	rt "runtime"
)

func RegisterAll() {
	p := false

	if rt.GOOS != "windows" {
		p = true
	}

	kFS.Register()
	kOS.Register(p)
	kPATH.Register(p)
}
