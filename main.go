package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	_ "www.lffwl.com/internal/packed"

	"www.lffwl.com/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
