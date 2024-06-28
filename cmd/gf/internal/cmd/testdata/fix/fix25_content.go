package testdata

import (
	"github.com/joy12825/gf/v2/frame/g"
	"github.com/joy12825/gf/v2/net/ghttp"
	"github.com/joy12825/gf/v2/util/guid"
	"testing"
)

func Test_Router_Hook_Multi(t *testing.T) {
	s := g.Server(guid.S())
	s.BindHandler("/multi-hook", func(r *ghttp.Request) {
		r.Response.Write("show")
	})

	s.BindHookHandlerByMap("/multi-hook", map[string]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			r.Response.Write("1")
		},
	})
	s.BindHookHandlerByMap("/multi-hook/{id}", map[string]ghttp.HandlerFunc{
		ghttp.HookBeforeServe: func(r *ghttp.Request) {
			r.Response.Write("2")
		},
	})
}
