//go:build tools
// +build tools

// goimports等の、ソースではimportしないけど、プロジェクトで利用するツールを管理するためのファイル 通常のビルドには含めない
// ここで、ツール類を定義しておけば、go mod tidy しても、go.mod から消されないし、バージョンが固定できる
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "golang.org/x/tools/cmd/goimports"
)
