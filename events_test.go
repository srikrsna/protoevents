package protoevents_test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/srikrsna/protoevents"
	expb "github.com/srikrsna/protoevents/example"
	"github.com/srikrsna/protoevents/mocks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestInterceptor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	disp := mocks.NewMockDispatcher(ctrl)
	disp.EXPECT().Dispatch(gomock.Any(), gomock.Any()).Times(1)

	el := mocks.NewMockErrorLogger(ctrl)
	el.EXPECT().Log(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)

	i := protoevents.NewInterceptor(protoevents.Options{
		Dispatcher:  disp,
		ErrorLogger: el,
	})

	lis := bufconn.Listen(1024 * 1024)

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(i))
	expb.RegisterExampleServiceServer(s, exampleService{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	bd := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bd), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := expb.NewExampleServiceClient(conn)
	client.ExampleFiringRpc(ctx, &expb.ExampleRpcRequest{})
	client.ExampleSilentRpc(ctx, &expb.ExampleRpcRequest{})
}

type exampleService struct {
	expb.UnimplementedExampleServiceServer
}

func (exampleService) ExampleFiringRpc(context.Context, *expb.ExampleRpcRequest) (*expb.ExampleRpcResponse, error) {
	return &expb.ExampleRpcResponse{}, nil
}
func (exampleService) ExampleSilentRpc(context.Context, *expb.ExampleRpcRequest) (*expb.ExampleRpcResponse, error) {
	return &expb.ExampleRpcResponse{}, nil
}
