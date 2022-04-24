# the 'fs' module

- [copyFile](#copyFile)
- [readFile](#readFile)
- [writeFile](#writeFile)
- [appendFile](#appendFile)
- [exists](#exists)

## copyFile

usage:

```js
const fs = require("fs")

fs.copyFile("example.txt", "newexample.txt")
```

## readFile

usage:

```js
const fs = require("fs")

const data = fs.readFile("example.txt")

// data => ["line 1", "line 2", "line 3...
```

## writeFile

usage:

```js
const fs = require("fs")

fs.writeFile("example.txt", [
	"line 1",
	"line 2",
	"line 3",
])
```

## appendFile

usage:

```js
const fs = require("fs")

fs.appendFile("example.txt", [
	"line 1",
	"line 2",
	"line 3",
])
```

## exists

usage:

```js
const fs = require("fs")

if (fs.exists("example.txt")) {
	console.log("file or directory exists")
} else {
	console.log("file or directory exists")
}
```