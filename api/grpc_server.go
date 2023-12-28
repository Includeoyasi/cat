package api

import (
	"context"
	grpc "github.com/Includeoyasi/cat/pkg/cat"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type GrpcServer struct {
	grpc.UnimplementedCatServer
}

func GetCapsule(ctx context.Context, in *grpc.GetCapsuleRequest) (*grpc.GetCapsuleResponce, error) {
	log.Printf("Capsule ID: %s", in.GetId())
	return &grpc.GetCapsuleResponce{
		Result: &grpc.Capsule{
			Id:        505,
			Author:    "Dimka",
			Text:      "Hello world",
			Label:     "test",
			CreatedAt: timestamppb.New(time.Now()),
		},
	}, nil
}
