# the 'fs' module

- [copyFile](#copyFile)
- [readFile](#readFile)
- [writeFile](#writeFile)
- [appendFile](#appendFile)
- [exists](#exists)

## copyFile

usage:

```js
copyFile("example.txt", "newexample.txt")
```

## readFile

usage:

```js
const data = readFile("example.txt")

// data => ["line 1", "line 2", "line 3...
```

## writeFile

usage:

```js
writeFile("example.txt", [
	"line 1",
	"line 2",
	"line 3",
])
```

## appendFile

usage:

```js
appendFile("example.txt", [
	"line 1",
	"line 2",
	"line 3",
])
```

## exists

usage:

```js
if (exists("example.txt")) {
	console.log("file or directory exists")
} else {
	console.log("file or directory exists")
}
```