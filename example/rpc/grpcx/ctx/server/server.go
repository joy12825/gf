// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/joy12825/gf.

package main

import (
	"github.com/joy12825/gf/example/rpc/grpcx/ctx/controller"
)

func main() {
	s := grpcx.Server.New()
	controller.Register(s)
	s.Run()
}
