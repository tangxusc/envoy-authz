package server

import (
	"context"
	"fmt"
	envoy_service_auth_v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"google.golang.org/grpc"
	"testing"
)

func TestGrpcClient(t *testing.T) {
	dial, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer dial.Close()
	client := envoy_service_auth_v2.NewAuthorizationClient(dial)
	check, err := client.Check(context.TODO(), &envoy_service_auth_v2.CheckRequest{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(check)
}
