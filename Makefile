.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run-user-rpc
run-user-rpc:
	cd user-rpc && go run user.go

.PHONY: user-rpc
user-rpc:
	cd user-rpc && goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=. -m

.PHONY: run-redeem-code-rpc
run-redeem-code-rpc:
	cd redeem-code-rpc && go run redeemcode.go

.PHONY: redeem-code-rpc
redeem-code-rpc:
	cd redeem-code-rpc/pb && goctl rpc protoc redeem_code.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../

help:
	@echo "make help - 显示帮助信息"
	@echo "make tidy - go mod tidy"
	@echo "make run-user-rpc - 运行 user-rpc 服务"
	@echo "make user-rpc - 从 .proto 文件生成 Go 代码"
	@echo "make run-redeem-code-rpc - 运行 redeem-code-rpc 服务"
	@echo "make redeem-code-rpc - 从 .proto 文件生成 Go 代码"

.DEFAULT_GOAL := help
