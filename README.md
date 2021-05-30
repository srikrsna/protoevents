# protoevents

protoevents is a protoc extension used to dispatch events upon successful completion of grpc method calls. As of now it only supports Go.

## Usecases

It enables loosely coupled event driven architecture for various background tasks such as,
* sending emails
* sending sms
* sending notifications
* collecting payments
* trigerring integrations

## Usage

For a protobuf definition like this,

```proto
syntax = "proto3";

package example;

import "events/events.proto";

option go_package = "github.com/srikrsna/protoevents/example;expb";

service ExampleService {
  rpc ExampleFiringRpc(ExampleRpcRequest) returns (ExampleRpcResponse) {
    option (events.fire) = true;
  }

  rpc ExampleSilentRpc(ExampleRpcRequest) returns (ExampleRpcResponse);
}

message ExampleRpcRequest {}

message ExampleRpcResponse {}
```

```go
package main

func main() {
    grpc.NewServer(grpc.ChainUnaryInterceptor(
        protoevents.NewInterceptor(protoevents.Options{
            Dispatcher:  // TODO,
            ErrorLogger: // TODO,
	    },
    )))
}
```

Upon each succesful completion of the rpc tagged with `(events.fires) = true` an event will be dispatched with request, response, and full method name.

The dispatcher is configurable. Currently it supports the following dispatchers and encoding formats


Dispatchers:

[gocloud.dev/pubsub](https://pkg.go.dev/github.com/srikrsna/protoevents/dispatchers/gocloud)

Encoding:

[protobuf](https://pkg.go.dev/github.com/srikrsna/protoevents/encoding/proto)


