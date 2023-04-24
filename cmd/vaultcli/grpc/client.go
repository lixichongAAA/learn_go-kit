package grpc

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	vault "github.com/lixichongAAA/gokitexample"
	"github.com/lixichongAAA/gokitexample/pb"
	"google.golang.org/grpc"
)

// New make a new gokitexample.Service client.
func New(conn *grpc.ClientConn) vault.Service {
	var hashEndpoint = grpctransport.NewClient(
		conn, "Vault", "Hash",
		vault.EncodeGRPCHashRequest,
		vault.DecodeGRPCHashResponse,
		pb.HashResponse{},
	).Endpoint()
	var validateEndpoint = grpctransport.NewClient(
		conn, "Vault", "Validate",
		vault.EncodeGRPCValidateRequest,
		vault.DecodeGRPCHashResponse,
		pb.ValidateResponse{},
	).Endpoint()

	return vault.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}
}

// func NewTokenBucketLimitter(tb *rate.Limiter) endpoint.Middleware {
// 	return func(e endpoint.Endpoint) endpoint.Endpoint {
// 		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
// 			if !tb.Allow() {
// 				return nil, errors.New("Rate limit execced!")
// 			}
// 			return e(ctx, request)
// 		}
// 	}
// }
