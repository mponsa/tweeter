package main

import (
	"github.com/mponsa/tweeter/src/console"
	"github.com/mponsa/tweeter/src/rest"
)

func main() {
	go console.Run()
	rest.Run()
}
