package main

import (
	"flag"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/ashniu123/thrift-graphql-demo/service-user/cmd/impl"
	"github.com/ashniu123/thrift-graphql-demo/service-user/cmd/server"
	"github.com/ashniu123/thrift-graphql-demo/service-user/internal/thrift-rpc/user"
)

var (
	addr = flag.String("addr", "localhost:9090", "Thrift Address to listen on")
)

var ApiVersion = 1

func main() {
	flag.Parse()

	impl.ApiVersion = user.Int(ApiVersion)

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	if err := server.RunServer(transportFactory, protocolFactory, *addr, false); err != nil {
		log.Println("error running thrift server: ", err)
	}
}
