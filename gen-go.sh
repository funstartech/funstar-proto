protoc --go_out=go --proto_path=proto proto/*.proto
protoc --go-grpc_out=go --proto_path=proto proto/*.proto
protoc --grpc-gateway_out=grpc_api_configuration=gateway-rules.yaml:go --proto_path=proto proto/*.proto
