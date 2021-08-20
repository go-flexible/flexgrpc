package flexgrpc_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/go-flexible/flexgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func TestHealthCheck(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := flexgrpc.New(&flexgrpc.Config{
		Addr: ":50051",
	})
	go func() {
		_ = server.Run(ctx)
	}()

	// give it time to start.
	time.Sleep(time.Second)

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	client := grpc_health_v1.NewHealthClient(conn)
	res, err := client.Check(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: "",
	})
	if err != nil {
		t.Error(err)
	}
	equal(t, "SERVING", res.Status.String())
}

func Example() {
	srv := flexgrpc.New(
		&flexgrpc.Config{Addr: ":8080"},
		grpc.ConnectionTimeout(10*time.Second),
	)
	_ = srv.Run(context.Background())
}

func equal(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %#[1]v (%[1]T), but wanted: %#[2]v (%[2]T)", got, want)
	}
}
