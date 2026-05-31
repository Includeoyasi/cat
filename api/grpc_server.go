package api

import (
	"context"
	"log"
	"time"

	catpb "github.com/Includeoyasi/cat/pkg/cat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	catpb.UnimplementedCatServer
	capsules map[int32]*catpb.Capsule
}

func NewGrpcServer() GrpcServer {
	createdAt := timestamppb.New(time.Now())

	return GrpcServer{
		capsules: map[int32]*catpb.Capsule{
			1: {
				Id:        1,
				Author:    "Includeoyasi",
				Text:      "First memory capsule",
				Label:     "demo",
				CreatedAt: createdAt,
			},
		},
	}
}

func (s GrpcServer) GetCapsule(ctx context.Context, in *catpb.GetCapsuleRequest) (*catpb.GetCapsuleResponce, error) {
	log.Printf("Capsule ID: %d", in.GetId())

	capsule, ok := s.capsules[in.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "capsule %d not found", in.GetId())
	}

	return &catpb.GetCapsuleResponce{
		Result: capsule,
	}, nil
}
