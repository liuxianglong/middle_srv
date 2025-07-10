package main

import (
	"middle_srv/app/rpc/internal/cmd"
	_ "middle_srv/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
