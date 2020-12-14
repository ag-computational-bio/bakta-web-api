protoc --proto_path=. --go-grpc_out=./proto/build --go_out=./proto/build --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/api/*.proto
protoc -I . --grpc-gateway_out ./proto/build --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative proto/api/*

protoc -I . --openapiv2_out ./proto/build --openapiv2_opt logtostderr=true proto/api/*

mv proto/build/proto/api/*go go/api
mv proto/build/proto/api/Jobs.swagger.json go/swaggerhandler

go generate go/swaggerhandler/SwaggerGen.go