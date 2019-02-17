package main

//go:generate simpleassets -a "./*js" -o assets.go
import "fmt"

func main() {
	base64, err := ReadAsset("base64.js")
	if err != nil {
		panic(err)
	}
	fmt.Println(base64)
}
