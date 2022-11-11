SHELL=/bin/bash

GOMOD:=GO111MODULE=on
export GOBIN := $(PWD)/bin
GO_PATH:=$(shell go env GOPATH)
export PATH:=$(GOBIN):$(GO_PATH):$(PATH)

## toolをinstallする
install-tools:
	@$(GOMOD) /bin/bash ./tools/install_tools.sh ./tools/tools.go

## 依存するpackageを最新版に更新する。
update-module:
	$(GOMOD) go mod tidy

## .go-versonに記載されているバージョンのgoをインストール
install-go:
	@./tools/install_go.sh

## 最初の準備するために全てのpackagesをinstallする
bootstrap: install-go install-tools update-module
