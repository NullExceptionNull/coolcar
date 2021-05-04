 protoc --go_out=./gen/go --go_opt=paths=source_relative \
    --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative \
   auth.proto

protoc -I . --grpc-gateway_out ./gen/go \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    auth.proto