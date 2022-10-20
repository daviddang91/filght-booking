grpc-customer-gen:
    protoc customer/grpc/proto/*.proto --go-grpc_out=.
    protoc customer/grpc/proto/*.proto --go_out=.