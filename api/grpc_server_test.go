package api

import (
	"context"
	"testing"

	catpb "github.com/Includeoyasi/cat/pkg/cat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetCapsuleReturnsSeededCapsule(t *testing.T) {
	server := NewGrpcServer()

	res, err := server.GetCapsule(context.Background(), &catpb.GetCapsuleRequest{Id: 1})
	if err != nil {
		t.Fatalf("get capsule: %v", err)
	}

	if res.GetResult().GetId() != 1 {
		t.Fatalf("unexpected capsule id: got %d, want 1", res.GetResult().GetId())
	}
}

func TestGetCapsuleReturnsNotFound(t *testing.T) {
	server := NewGrpcServer()

	_, err := server.GetCapsule(context.Background(), &catpb.GetCapsuleRequest{Id: 404})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("unexpected error code: got %v, want %v", status.Code(err), codes.NotFound)
	}
}
