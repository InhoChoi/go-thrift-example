package main

import (
  "git.apache.org/thrift.git/lib/go/thrift"
  "fmt"
)

func main()  {
  var addr string = "localhost:9090"
  var secure bool = false
  var protocolFactory thrift.TProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
  var transportFactory thrift.TTransportFactory = thrift.NewTTransportFactory()

  if err := runClient(transportFactory, protocolFactory, addr, secure); err != nil {
    fmt.Println("error running client:", err)
  }
}