grpc-gen:
    protoc common/grpc/proto/*.proto --go-grpc_out=.
    protoc customer/grpc/proto/*.proto --go_out=.

run-customer:
	CompileDaemon --build="go build -o ./out/customer customer/main.go" --command="./out/customer"

run-flight:
	CompileDaemon --build="go build -o ./out/flight flight/main.go" --command="./out/flight"

run-booking:
	CompileDaemon --build="go build -o ./out/booking booking/main.go" --command="./out/booking"