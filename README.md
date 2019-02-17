# simpleassets

This is a dead simple code generator for static assets. It reads, compresses and encodes files into Go for runtime access. The files are encoded with base64 and stored in a map. The files are decoded and decompressed when accessed. Additional assets can be added during runtime but are only encoded, not compressed.

Files are added via glob patterns. Multiple patterns are supported. The additional write functionality can be disabled during the generation. Output is a single Go file that has beem formatted with `gofmt`.

## Install

```
go get github.com/eliquious/simpleassets
```

## Usage

```
Usage of simpleassets:
  -a, --assets strings   Asset glob(s)
      --omit-write       Omit the WriteAsset generated function
  -o, --output string    Output file (default "assets.go")
  -p, --package string   Package name (default "main")
  -t, --trim string      Trim file prefix
```

### Features

- Multiple glob patterns
- Single output
- Gzip compression
- Supports runtime additions
- Trim file prefix

## Example

This code is provided in the examples directory. Make sure to run `go generate` from within the examples folder before running. 

```go
package main

//go:generate simpleassets -a "./*js" -o assets.go
import "fmt"

func main() {

	// base64.js is a static asset produced from `go generate`
	base64, err := ReadAsset("base64.js")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(base64))

	// Write temporary data into asset storage.
	WriteAsset("1990s.txt", []byte("All your base are belong to us"))
}
```
