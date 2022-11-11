
## install_tools.sh と tools.go
tools.goには、toolとして利用したいものが書かれている。例えば、goimportsとか
install.shを読み込むことでtools.goの中身を読み取る。

どうしてtools.goに使うツールを指定してるかというと、バージョンを固定するため。
tools.goにツールを書いて、go mod tidyをすることでgo.modにツールのバージョンが記載される。
go install する際にmod.goのバージョンを参照してgoimportsなどはインストールされる。

## golandでinstallしたgoimportsを利用するためには

```shell
$ make install-tools
```

binにファイルが保存される。
golandのfile watcherで,goimportsを作成. Programを`$ProjectFileDir$/bin/goimports`に設定する。


### tools.goのコメントについて
上のコメントはbuild constrainというもの。簡単にいうとbuildするときにtagでtoolsを指定しないと、tools.goはbuildに取り込まれない。
build constrainで調べると色々でてくる。

下記のように２つの表記があるのはバージョンによって書き方が異なるため。[参考](https://future-architect.github.io/articles/20210810a/#build-constraint%E3%81%AE%E3%82%B3%E3%83%A1%E3%83%B3%E3%83%88%E3%81%8C%E5%A4%89%E6%9B%B4)

//go:build tools  
// +build tools

