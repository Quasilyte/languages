package main

import (
	"fmt"

	"github.com/Quasilyte/languages"
)

func main() {
	for _, l := range languages.List() {
		fmt.Println(l.Name())
	}
}
