protoc --proto_path=. --go-grpc_out=./proto/build --go_out=./proto/build --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/api/*.proto

mv proto/build/proto/api/*go go/api