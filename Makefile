.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run-user-rpc
run-user-rpc:
	cd user-rpc && go run user.go

.PHONY: rpc-user
rpc-user:
	cd user-rpc && goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m

help:
	@echo "make help - 显示帮助信息"
	@echo "make tidy - go mod tidy"
	@echo "make run-user-rpc - 运行 user-rpc 服务"
	@echo "make rpc-user - 从 .proto 文件生成 Go 代码"

.DEFAULT_GOAL := help
