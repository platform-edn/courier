package server

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/platform-edn/courier/mock"
	"github.com/platform-edn/courier/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestMessageServer_PublishMessage(t *testing.T) {
	server := NewMessageServer()
	push := server.PushChannel()

	startTestServer(server)

	client, conn, err := mock.NewLocalGRPCClient("bufnet", bufDialer)
	if err != nil {
		t.Fatalf("could not create client: %s", err)
	}
	defer conn.Close()

	go func() {
		ctx := context.Background()
		m := proto.PublishMessageRequest{
			Message: &proto.PublishMessage{
				Id:      uuid.NewString(),
				Subject: "test",
				Content: []byte("test"),
			},
		}

		client.PublishMessage(ctx, &m)
	}()

	select {
	case <-push:
		break
	case <-time.After(time.Second * 1):
		t.Fatal("timeout waitng on push channel")
	}
}

func TestMessageServer_ResponseMessage(t *testing.T) {
	server := NewMessageServer()
	push := server.PushChannel()

	startTestServer(server)

	client, conn, err := mock.NewLocalGRPCClient("bufnet", bufDialer)
	if err != nil {
		t.Fatalf("could not create client: %s", err)
	}
	defer conn.Close()

	go func() {
		ctx := context.Background()
		m := proto.ResponseMessageRequest{
			Message: &proto.ResponseMessage{
				Id:      uuid.NewString(),
				Subject: "test",
				Content: []byte("test"),
			},
		}

		client.ResponseMessage(ctx, &m)
	}()

	select {
	case <-push:
		break
	case <-time.After(time.Second * 1):
		t.Fatal("timeout waitng on push channel")
	}
}

func TestMessageServer_RequestMessage(t *testing.T) {
	server := NewMessageServer()
	push := server.PushChannel()
	info := server.ResponseChannel()

	startTestServer(server)
	client, conn, err := mock.NewLocalGRPCClient("bufnet", bufDialer)
	if err != nil {
		t.Fatalf("could not create client: %s", err)
	}
	defer conn.Close()

	go func() {
		ctx := context.Background()
		m := proto.RequestMessageRequest{
			Message: &proto.RequestMessage{
				Id:      uuid.NewString(),
				NodeId:  uuid.NewString(),
				Subject: "test",
				Content: []byte("test"),
			},
		}

		client.RequestMessage(ctx, &m)
	}()

	p := false
	i := false
	r := 0
channel:
	for {
		select {
		case <-push:
			p = true
		case <-info:
			i = true
		default:
			if p && i {
				break channel
			}
			if r == 5 {
				t.Fatalf("failed waiting for info and push channels")
			}
			time.Sleep(time.Millisecond * 100)
			r++
		}
	}
}

func startTestServer(m *MessageServer) {
	lis = bufconn.Listen(bufSize)

	grpcServer := grpc.NewServer()

	proto.RegisterMessageServerServer(grpcServer, m)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}