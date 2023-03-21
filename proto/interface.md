# gRPC Service Interface

This file serves as an interface that must be implemented by all subapps. All subapps must support the rpc call as described below

```proto
rpc Describe (google.protobuf.Empty) returns (pichan.Description)
```
