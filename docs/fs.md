# the `fs` module

- [appendFile](#appendFile)
- [writeFile](#writeFile)
- [readFile](#readFile)
- [rename](#rename)
- [copy](#copy)
- [remove](#remove)
- [isDir](#isDir)
- [mkdir](#mkdir)
- [chmod](#chmod)
- [chown](#chown)
- [exists](#exists)

## appendFile

```js
const fs = require("fs")

fs.appendFile("example.txt", [
	"line 1",
	"line 2",
	"line 3",
])
```

## writeFile

```js
const fs = require("fs")

fs.writeFile("example.txt", [
	"line 1",
	"line 2",
	"line 3",
])
```

## readFile

```js
const fs = require("fs")

const data = fs.readFile("example.txt")

// data => ["line 1", "line 2", "line 3...
```

## rename

```js
const fs = require("fs")

fs.rename("example.txt", "newexample.txt")
```

## copy

```js
const fs = require("fs")

fs.copy("example.txt", "newexample.txt")
```

## remove

```js
const fs = require("fs")

fs.remove("example.txt")
```

## isDir

```js
const fs = require("fs")

if (fs.isDir("exampledir")) {
	console.log("path is a directory")
}
```

## mkdir

```js
const fs = require("fs")

fs.mkdir("exampledir")
```

## chmod

```js
const fs = require("fs")

fs.chmod("example.txt", 0666)
```

## chown

```js
const fs = require("fs")

fs.chown("example.txt", 1000, 1000)
```

## exists

```js
const fs = require("fs")

if (fs.exists("example.txt")) {
	console.log("file or directory exists")
} else {
	console.log("file or directory exists")
}
```