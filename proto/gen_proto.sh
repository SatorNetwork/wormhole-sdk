protoc -I . \
    -I $GOPATH/src/github.com/googleapis/googleapis \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    publicrpc/v1/publicrpc.proto \
    gossip/v1/gossip.proto \
    node/v1/node.proto