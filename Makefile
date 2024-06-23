RPCDIR=$(CURDIR)/rpc
OUTDIR=$(CURDIR)/dist
GO_AUTOGENED=$(CURDIR)/pkg/autogen/api
GO_SRVBIN=$(OUTDIR)/planb

# ビルド
.PHONY: all
all: $(GO_AUTOGENED) $(GO_SRVBIN)

# protocol bufferのビルド
.PHONY: bufbuild
bufbuild: $(GO_AUTOGENED)

# サーバーのビルド
$(GO_SRVBIN): $(shell find $(CURDIR) -type f -name '*.go' -print)
	gofmt -l -w $?
	staticcheck ./...
	go vet ./...
	go build -o $@ $(GO_SRVDIR)

# buf側のビルド前処理
$(GO_AUTOGENED): $(shell find $(RPCDIR) -type f -name '*.proto' -print)
	buf dep update $(RPCDIR)
	buf lint
	buf format -w
	buf generate

# 実行ファイルの削除
.PHONY: clean
clean:
	@rm -rf $(OUTDIR)/* $(CURDIR)/pkg/autogen/