package std

import (
	kPROCESS "github.com/kosmosJS/std/lib/process"
	kPATH "github.com/kosmosJS/std/lib/path"
	kMEM "github.com/kosmosJS/std/lib/mem"
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
	kMEM.Register()
	kPROCESS.Register()
}
