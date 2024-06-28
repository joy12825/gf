package article

import (
	"context"

	"github.com/joy12825/gf/v2/errors/gcode"
	"github.com/joy12825/gf/v2/errors/gerror"

	"github.com/joy12825/gf/cmd/gf/v2/internal/cmd/testdata/genctrl/api/article/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
