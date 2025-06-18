.PHONY: generator storage user

generator: generator/generator.proto
	protoc \
	--go_out=generator \
	--go-grpc_out=generator \
	--go_opt=module=github.com/child6yo/rago/proto/generator \
	--go-grpc_opt=module=github.com/child6yo/rago/proto/generator \
	--proto_path=. \
	generator/generator.proto

storage: storage/storage.proto
	protoc \
	--go_out=storage \
	--go-grpc_out=storage \
	--go_opt=module=github.com/child6yo/rago/proto/storage \
	--go-grpc_opt=module=github.com/child6yo/rago/proto/storage \
	--proto_path=. \
	storage/storage.proto

user: user/user.proto
	protoc \
	--go_out=user \
	--go-grpc_out=user \
	--go_opt=module=github.com/child6yo/rago/proto/user \
	--go-grpc_opt=module=github.com/child6yo/rago/proto/user \
	--proto_path=. \
	user/user.proto

all: generator storage user