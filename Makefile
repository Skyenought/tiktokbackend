STYLE := -style=goZero

TEMPLATE := -home=pkg/tpl
REL_TEMPLATE := -home=../../../pkg/tpl

type ?= ""
service ?= ""

.PHONY: rpc
rpc:
	cp pkg/idl/$(type).proto cmd/$(type)/rpc
	cd cmd/$(type)/rpc && goctl rpc protoc $(type).proto \
	--go_out=. --go-grpc_out=. --zrpc_out=. $(STYLE) $(TEMPLATE)

.PHONY: api
api:
	goctl api go -api=pkg/api/$(type).api $(STYLE) -dir=cmd/$(type)/api $(TEMPLATE)

.PHONY: dev
dev:
	cd pkg/scripts && .\develop.bat -t $(type) -s $(service)

.PHONY: rm
rm:
	cd cmd/$(type)/$(service) && rm -rf *

.PHONY: fmt
fmt:
	gofumpt -l -w .                                                                                                                                                ─╯

modd:
	modd

