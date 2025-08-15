package frontend

import (
	"context"
	"fmt"
	baseApi "gf-chat/api"
	v1 "gf-chat/api/frontend/v1"
	"gf-chat/internal/service"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
)

var (
	update = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var CWs = &cWs{}

type cWs struct {
}

func (c cWs) Connect(ctx context.Context, _ *v1.ChatConnectReq) (res *baseApi.NilRes, err error) {
	request := ghttp.RequestFromCtx(ctx)
	conn, err := update.Upgrade(request.Response.Writer, request.Request, nil)
	if err != nil {
		request.Exit()
		return
	}
	user := service.UserCtx().GetUser(ctx)
	fmt.Printf("ws User: %v\n", user)
	err = service.Chat().Register(ctx, user, conn, service.Platform().GetPlatform(ctx))
	if err != nil {
		return
	}
	res = baseApi.NewNilResp()
	return
}
