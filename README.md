# cat

cat is a small gRPC service for managing memory capsules.

It exposes a `Cat` gRPC service with methods for creating, reading, and deleting capsules. The current implementation uses an in-memory store, which keeps the service simple while the API contract is being shaped.

## Run

```sh
go run .
```

The gRPC server listens on port `50051`.

## API

```proto
service Cat {
  rpc GetCapsule(GetCapsuleRequest) returns (GetCapsuleResponce);
  rpc CreateCapsule(CreateCapsuleRequest) returns (CreateCapsuleResponce);
  rpc DeleteCapsule(DeleteCapsuleRequest) returns (DeleteCapsuleResponce);
}
```

## Development

Regenerate protobuf files:

```sh
make proto
```

Run checks:

```sh
go test ./...
```
