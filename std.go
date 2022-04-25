package std

import (
	kFS "github.com/kosmosJS/std/fs"
	kOS "github.com/kosmosJS/std/os"
)

func RegisterAll() {
	kFS.Register()
	kOS.Register()
}
