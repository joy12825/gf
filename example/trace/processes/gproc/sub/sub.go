// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/joy12825/gf.

package main

import (
	"github.com/joy12825/gf/v2/frame/g"
	"github.com/joy12825/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.GetInitCtx()
	g.Log().Debug(ctx, `this is sub process`)
}
