// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/joy12825/gf.

package main

import (
	_ "github.com/joy12825/gf/example/config/polaris/boot"

	"github.com/joy12825/gf/v2/frame/g"
	"github.com/joy12825/gf/v2/os/gctx"
)

func main() {
	var ctx = gctx.GetInitCtx()

	// Available checks.
	g.Dump(g.Cfg().Available(ctx))

	// All key-value configurations.
	g.Dump(g.Cfg().Data(ctx))

	// Retrieve certain value by key.
	g.Dump(g.Cfg().MustGet(ctx, "server.address"))
}
