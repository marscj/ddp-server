package main

import (
	_ "ddp-server/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"ddp-server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
