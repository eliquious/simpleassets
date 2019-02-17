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
	WriteAsset("1990s.txt", []byte("All ur base are belong to us"))
}
