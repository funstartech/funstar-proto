protoc --go_out=go proto/*.proto
protoc --go-grpc_out=go proto/*.proto
protoc --grpc-gateway_out=grpc_api_configuration=gateway-rules.yaml:go proto/*.proto
