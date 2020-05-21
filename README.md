# goによるTodoアプリ実装
# 学習要素
goによる以下項目の実装
* HTTPリクエストハンドラ（gin, http）
* データベースとの連携
* セッションマネージャ

## パッケージ
* gin

## Intellij
go modules

## コマンドメモ
```bash
go mod init go_todo_sample

go get github.com/gin-gonic/gin
go run main.go

# realize（ホットリロード）の導入 https://github.com/oxequa/realize/issues/253
# GO111MODULE=off go get github.com/oxequa/realize 解決しない


```

## realize（ホットリロード）の導入
* https://github.com/oxequa/realize/issues/253
    * GO111MODULE=off go get github.com/oxequa/realize 解決しない（go modules以前のやり方）
```bash
# go.modに以下追記
replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.1.1

# その後、以下コマンド
go get -u github.com/oxequa/realize

# .realize.yamlの編集後、
realize start
```

## 参考文献
* [Webアプリ初心者がGo言語でサーバサイド（1. 簡単なHTTPサーバの実装）](https://qiita.com/wsuzume/items/75d5c0cd2dd5a1963b9e)
* [Go Web プログラミング](https://astaxie.gitbooks.io/build-web-application-with-golang/ja/)
* [Go Doc](https://godoc.org/?q=main)

* [【Go言語】パッケージのimportについて整理した](https://qiita.com/ogady/items/0cedd3599c4dc13e9a95)
