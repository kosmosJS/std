# the `fs` module

- [copyFile](#copyFile)
- [readFile](#readFile)
- [writeFile](#writeFile)
- [appendFile](#appendFile)
- [exists](#exists)

## copyFile

```js
const fs = require("fs")

fs.copyFile("example.txt", "newexample.txt")
```

## readFile

```js
const fs = require("fs")

const data = fs.readFile("example.txt")

// data => ["line 1", "line 2", "line 3...
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

## appendFile

```js
const fs = require("fs")

fs.appendFile("example.txt", [
	"line 1",
	"line 2",
	"line 3",
])
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