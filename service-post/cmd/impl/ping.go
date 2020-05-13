package impl

import (
	"context"
	"log"

	"github.com/ashniu123/thrift-graphql-demo/service-post/internal/thrift-rpc/post"
)

var ApiVersion post.Int

func (p *PostHandler) Ping(ctx context.Context) (*post.PingResponse, error) {
	message := "ping"
	log.Println(message)

	PingResponse := post.NewPingResponse()
	PingResponse.Message = &message
	PingResponse.Version = &ApiVersion

	return PingResponse, nil
}
