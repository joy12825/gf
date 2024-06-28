// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/joy12825/gf.

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/polarismesh/polaris-go/api"
	"github.com/polarismesh/polaris-go/pkg/config"

	"github.com/joy12825/gf/v2/frame/g"
	"github.com/joy12825/gf/v2/net/gsvc"
	"github.com/joy12825/gf/v2/os/gctx"
)

func main() {
	conf := config.NewDefaultConfiguration([]string{"183.47.111.80:8091"})
	conf.Consumer.LocalCache.SetPersistDir("/tmp/polaris/backup")
	if err := api.SetLoggersDir("/tmp/polaris/log"); err != nil {
		g.Log().Fatal(context.Background(), err)
	}

	gsvc.SetRegistry(polaris.NewWithConfig(conf, polaris.WithTTL(10)))

	for i := 0; i < 100; i++ {
		res, err := g.Client().Get(gctx.New(), `http://hello-world.svc/`)
		if err != nil {
			panic(err)
		}
		fmt.Println(res.ReadAllString())
		res.Close()
		time.Sleep(time.Second)
	}
}
