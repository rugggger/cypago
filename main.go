package main

import (
	"fmt"
	"github.com/cypago/pkg/app"
)

func main() {
	fmt.Println("hi")
	scanApp := app.New()
	scanApp.Serve()

}
