package main

import (
	_ "shop/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
