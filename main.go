package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	svr := server.Default()

	svr.GET("/ping", func(ctx context.Context, req *app.RequestContext) {
		req.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	svr.Spin()
}
