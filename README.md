# flexgrpc

A flex compatible grpc server

The package `flexgrpc` provides a default set of configuration for hosting a grpc server in a service.

## Configuration

The GRPC server can be configured through these environment variables:

- `GRPC_ADDR` the gRPC server listener's network address (default: `0.0.0.0:50051`)

## Examples

### Starting server and exposing the service

```go
srv := flexgrpc.New(
    &flexgrpc.Config{Addr: ":8080"},
    grpc.ConnectionTimeout(10*time.Second),
)
_ = srv.Run(ctx)
```
