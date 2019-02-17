<p align="center"><img src="icon.png" width="180"></p>
<p align="center"><h1 class="">simpleassets</h1></p>

This is a dead simple code generator for static assets. It reads, compresses and encodes files into Go for runtime access. The files are encoded with base64 and stored in a map. The files are decoded and decompressed when accessed. Additional assets can be added during runtime but are only encoded, not compressed.

Files are added via glob patterns. Multiple patterns are supported. The additional write functionality can be disabled during generation. Output is a single Go file that is formatted with `gofmt`.

### Yet another... Why?!

Simply put, all the libraries I found were either too much or not enough. Many of the libraries out there are built for http servers or they don't have a way to write files to the memory store at run time. So `simpleasssets` is not built for serving assets and allows you to store new assets at runtime.

Everything I need and nothing I don't. It's a no-frills, just-load-the-damn-files type of tool. If that's what your looking for, awesome. Otherwise the larger, friller, non-sensical, swiss-army-knife serving tools might suit you better.

## Install

```
go get github.com/eliquious/simpleassets
```

#### Dependencies

```
go get -u github.com/spf13/pflag
```

## Features

- Multiple glob patterns
- Single output
- Gzip compression
- Supports runtime additions
- Trim file prefix

## Usage

```
Usage: simpleassets [flags] ASSETS [ASSETS ...] 
      --omit-write       Disable write access (not generated)
  -o, --output string    Output file (default "assets.go")
  -p, --package string   Package name (default "main")
  -t, --trim string      Trim file prefix
```

## Example

This code is provided in the examples directory. Make sure to run `go generate` from within the examples folder before running. 

```go
package main

//go:generate simpleassets -o assets.go ./*js
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
