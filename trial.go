package main

import (
	"fmt"
)

// Object declaration in Go
type answer struct {
	ID       string `json:"id"`
	USERNAME string `json:"username"`
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}

// Object initialization in Go
var album = []answer{
	{
		ID:       "1",
		USERNAME: "Duncan",
		EMAIL:    "duncanii414@gmail.com",
		PASSWORD: "password",
	},
}

func main() {
	fmt.Println(album)
}
