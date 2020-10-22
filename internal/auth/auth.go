package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	rpcpb "github.com/johnbellone/persona-service/internal/gen/persona/rpc"
	log "github.com/sirupsen/logrus"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Perform the reflection on the request to determine if method option is set.
		opts := req.(proto.Message).ProtoReflect().Descriptor().Options().(*descriptorpb.MethodOptions)
		access := proto.GetExtension(opts, rpcpb.E_MethodAccess).(rpcpb.MethodAccess)

		log.Warnf("%v", access)
		return handler(ctx, req)
	}
}
