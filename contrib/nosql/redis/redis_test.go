// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/joy12825/gf.

package redis_test

import (
	"github.com/joy12825/gf/v2/database/gredis"
	"github.com/joy12825/gf/v2/os/gctx"
)

var (
	ctx    = gctx.GetInitCtx()
	config = &gredis.Config{
		Address: `:6379`,
		Db:      1,
	}
	redis, _ = gredis.New(config)
)
