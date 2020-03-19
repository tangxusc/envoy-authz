package server

import (
	"context"
	"encoding/json"
	"fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/tangxusc/envoy-authz/pkg/config"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"net"
)

type AuthZServer struct {
}

func (a *AuthZServer) Check(ctx context.Context, ch *v2.CheckRequest) (cp *v2.CheckResponse, e error) {
	log.Println(">>> Authorization called check()")

	b, err := json.MarshalIndent(ch.Attributes.Request.Http.Headers, "", "  ")
	if err == nil {
		log.Println("Inbound Headers: ")
		log.Println(string(b))
	}

	ct, err := json.MarshalIndent(ch.Attributes.ContextExtensions, "", "  ")
	if err == nil {
		log.Println("Context Extensions: ")
		log.Println(string(ct))
	}
	//var cp *v2.CheckResponse

	options := make([]*core.HeaderValueOption, 1)
	options[0] = &core.HeaderValueOption{
		Header: &core.HeaderValue{
			Key:   "testAuthz",
			Value: "TestAuthZValue",
		},
		//Append: &wrappers.BoolValue{
		//	Value: true,
		//},
	}
	if config.Instance.Allow {
		okResponse := &v2.CheckResponse_OkResponse{
			OkResponse: &v2.OkHttpResponse{
				Headers: options,
			},
		}

		cp = &v2.CheckResponse{
			Status: &status.Status{
				Code:    int32(codes.OK),
				Message: "ok response message",
			},
			HttpResponse: okResponse,
		}
		fmt.Println("返回成功")
	} else {
		okResponse := &v2.CheckResponse_DeniedResponse{
			DeniedResponse: &v2.DeniedHttpResponse{
				Status: &envoy_type.HttpStatus{
					Code: envoy_type.StatusCode_Unauthorized,
				},
				Headers: options,
				Body:    "no authz",
			},
		}
		cp = &v2.CheckResponse{
			Status: &status.Status{
				Code:    int32(codes.Unauthenticated),
				Message: "cp message in this",
			},
			HttpResponse: okResponse,
		}
		fmt.Println("返回失败")
	}
	return cp, nil
}

func StartServer() error {
	server := grpc.NewServer()
	impl := &AuthZServer{}

	v2.RegisterAuthorizationServer(server, impl)
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		return err
	}
	err = server.Serve(listen)
	if err != nil {
		return err
	}
	defer server.GracefulStop()
	return nil
}
