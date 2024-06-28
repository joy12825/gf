// Copyright GoFrame gf Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/joy12825/gf.

package cmd

import (
	"context"

	"github.com/joy12825/gf/v2/frame/g"

	"github.com/joy12825/gf/cmd/gf/v2/internal/service"
)

var (
	Install = cInstall{}
)

type cInstall struct {
	g.Meta `name:"install" brief:"install gf binary to system (might need root/admin permission)"`
}

type cInstallInput struct {
	g.Meta `name:"install"`
}

type cInstallOutput struct{}

func (c cInstall) Index(ctx context.Context, in cInstallInput) (out *cInstallOutput, err error) {
	err = service.Install.Run(ctx)
	return
}
