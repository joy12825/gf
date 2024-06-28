package article

import (
	"context"

	"github.com/joy12825/gf/v2/errors/gcode"
	"github.com/joy12825/gf/v2/errors/gerror"

	"github.com/joy12825/gf/cmd/gf/v2/internal/cmd/testdata/genctrl/api/article/v2"
)

func (c *ControllerV2) Create(ctx context.Context, req *v2.CreateReq) (res *v2.CreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
