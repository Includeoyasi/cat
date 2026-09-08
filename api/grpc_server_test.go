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

func TestCreateCapsule(t *testing.T) {
	server := NewGrpcServer()

	res, err := server.CreateCapsule(context.Background(), &catpb.CreateCapsuleRequest{
		Author: "Includeoyasi",
		Text:   "Created from pet",
		Label:  "rest",
	})
	if err != nil {
		t.Fatalf("create capsule: %v", err)
	}

	if res.GetResult().GetId() != 2 {
		t.Fatalf("unexpected capsule id: got %d, want 2", res.GetResult().GetId())
	}

	got, err := server.GetCapsule(context.Background(), &catpb.GetCapsuleRequest{Id: res.GetResult().GetId()})
	if err != nil {
		t.Fatalf("get created capsule: %v", err)
	}
	if got.GetResult().GetText() != "Created from pet" {
		t.Fatalf("unexpected capsule text: got %q", got.GetResult().GetText())
	}
}

func TestDeleteCapsule(t *testing.T) {
	server := NewGrpcServer()

	res, err := server.DeleteCapsule(context.Background(), &catpb.DeleteCapsuleRequest{Id: 1})
	if err != nil {
		t.Fatalf("delete capsule: %v", err)
	}
	if !res.GetDeleted() {
		t.Fatal("delete response is false")
	}

	_, err = server.GetCapsule(context.Background(), &catpb.GetCapsuleRequest{Id: 1})
	if status.Code(err) != codes.NotFound {
		t.Fatalf("unexpected error code: got %v, want %v", status.Code(err), codes.NotFound)
	}
}
