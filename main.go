package main

import (
	_ "MySpace/boot"
	_ "MySpace/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
