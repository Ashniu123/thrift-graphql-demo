package server

import (
	"log"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/ashniu123/thrift-graphql-demo/service-user/cmd/impl"
	"github.com/ashniu123/thrift-graphql-demo/service-user/internal/thrift-rpc/user"
)

func RunServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	handler := impl.NewUserHandler()
	processor := user.NewUserServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	log.Printf("Starting simple thrift server on: %s\n", addr)
	return server.Serve()
}
