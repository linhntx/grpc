gen :
	protoc --go_out=.\grpc_chatserver --go_opt=paths=source_relative --go-grpc_out=.\grpc_chatserver --go-grpc_opt=paths=source_relative .\grpc_chat.proto