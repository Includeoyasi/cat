# cat

cat is a small gRPC service for reading memory capsules.

It exposes a `Cat` gRPC service with a `GetCapsule` method and returns capsule data by ID. The current implementation uses a small in-memory store, which keeps the service simple while the API contract is being shaped.

## Run

```sh
go run .
```

The gRPC server listens on port `50051`.

## API

```proto
service Cat {
  rpc GetCapsule(GetCapsuleRequest) returns (GetCapsuleResponce);
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
