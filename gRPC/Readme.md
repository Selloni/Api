brew install protoc-gen-go /
brew install protobuf 

protoc --go_out=pkg/api --go_opt=paths=source_relative \
--go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative \
proto/adder.proto


https://www.youtube.com/watch?v=z-mHhobE0Pw&t=868s