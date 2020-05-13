package impl

import (
	"context"
	"log"

	"github.com/ashniu123/thrift-graphql-demo/service-user/internal/thrift-rpc/user"
)

var ApiVersion user.Int

func (p *UserHandler) Ping(ctx context.Context) (*user.PingResponse, error) {
	message := "ping"
	log.Println(message)

	PingResponse := user.NewPingResponse()
	PingResponse.Message = &message
	PingResponse.Version = &ApiVersion

	return PingResponse, nil
}
