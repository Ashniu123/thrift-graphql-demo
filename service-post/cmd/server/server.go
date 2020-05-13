package server

import (
	"log"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/ashniu123/thrift-graphql-demo/service-post/cmd/impl"
	"github.com/ashniu123/thrift-graphql-demo/service-post/internal/thrift-rpc/post"
)

func RunServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	handler := impl.NewPostHandler()
	processor := post.NewPostServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	log.Printf("Starting simple thrift server on: %s\n", addr)
	return server.Serve()
}
