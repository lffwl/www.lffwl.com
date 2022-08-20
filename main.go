package main

import (
	_ "www.lffwl.com/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"www.lffwl.com/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
