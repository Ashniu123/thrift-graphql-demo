package main

import (
	"flag"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/ashniu123/thrift-graphql-demo/service-post/cmd/impl"
	"github.com/ashniu123/thrift-graphql-demo/service-post/cmd/server"
	"github.com/ashniu123/thrift-graphql-demo/service-post/internal/thrift-rpc/post"
)

var (
	addr = flag.String("addr", "localhost:9090", "Thrift Address to listen on")
)

var ApiVersion = 1

func main() {
	flag.Parse()
	impl.ApiVersion = post.Int(ApiVersion)

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	if err := server.RunServer(transportFactory, protocolFactory, *addr, false); err != nil {
		log.Println("error running thrift server: ", err)
	}
}
