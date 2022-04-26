package std

import (
	kPATH "github.com/kosmosJS/std/path"
	kFS "github.com/kosmosJS/std/fs"
	kOS "github.com/kosmosJS/std/os"
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
