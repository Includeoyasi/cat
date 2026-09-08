package api

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	catpb "github.com/Includeoyasi/cat/pkg/cat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcServer struct {
	catpb.UnimplementedCatServer
	mu       sync.RWMutex
	nextID   int32
	capsules map[int32]*catpb.Capsule
}

func NewGrpcServer() *GrpcServer {
	createdAt := timestamppb.New(time.Now())

	return &GrpcServer{
		nextID: 2,
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

	s.mu.RLock()
	defer s.mu.RUnlock()

	capsule, ok := s.capsules[in.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "capsule %d not found", in.GetId())
	}

	return &catpb.GetCapsuleResponce{
		Result: capsule,
	}, nil
}

func (s *GrpcServer) CreateCapsule(ctx context.Context, in *catpb.CreateCapsuleRequest) (*catpb.CreateCapsuleResponce, error) {
	author := strings.TrimSpace(in.GetAuthor())
	text := strings.TrimSpace(in.GetText())
	label := strings.TrimSpace(in.GetLabel())

	if author == "" {
		return nil, status.Error(codes.InvalidArgument, "author is required")
	}
	if text == "" {
		return nil, status.Error(codes.InvalidArgument, "text is required")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	capsule := &catpb.Capsule{
		Id:        s.nextID,
		Author:    author,
		Text:      text,
		Label:     label,
		CreatedAt: timestamppb.New(time.Now()),
	}

	s.capsules[capsule.Id] = capsule
	s.nextID++

	return &catpb.CreateCapsuleResponce{
		Result: capsule,
	}, nil
}

func (s *GrpcServer) DeleteCapsule(ctx context.Context, in *catpb.DeleteCapsuleRequest) (*catpb.DeleteCapsuleResponce, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.capsules[in.GetId()]; !ok {
		return nil, status.Errorf(codes.NotFound, "capsule %d not found", in.GetId())
	}

	delete(s.capsules, in.GetId())

	return &catpb.DeleteCapsuleResponce{
		Deleted: true,
	}, nil
}
