protogen: protogen-pichan protogen-habits

protogen-pichan:
	protoc -I=./pichan \
		--go_out=./pichan \
		--go_opt=paths=source_relative \
		--go-grpc_out=./pichan \
		--go-grpc_opt=paths=source_relative \
		./pichan/pichan.proto

protogen-habits:
	protoc -I=./pichan \
		-I=./habits \
		--go_out=./habits \
		--go_opt=paths=source_relative \
		--go-grpc_out=./habits \
		--go-grpc_opt=paths=source_relative \
		./habits/habits.proto
