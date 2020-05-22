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

## 基礎整理
### 先頭大文字(exported)
先頭大文字の関数、メソッド、定数、変数、構造体のフィールドは、外部パッケージから参照可能となる
名前に可視性が表現されている

### ポインタ
```go
var pointer *int // ポインタ変数（定義）
var n int = 100 // 値を入れている

pointer = &n // アドレス演算子（ポインタのアドレスを渡す）
fmt.Println("" + pointer) //-> アドレスの値
fmt.Println("" + *pointer) //-> 中身の値
```

```go
// ポインタによってアドレスを渡すことで、ほぼ参照渡しと同義のことができる（参照渡しは言語毎の仕組みにより安全性を考慮したポインタ渡しと同義）
func plusOne(a int, *b int) {
  a = a + 1 // 値渡し
  *b = *b + 1 // ポインタ渡し（参照渡し）
}

a, b := 10, 10
plusOne(a, b)

fmt.Println(a) // ->10
fmt.Println(b) // ->11
```

```go
// 先行したメモリの確保
var n *int = new(int)

type myStruct struct {
    a int
    b int
}
var myS *myStruct = new(myStruct)
```

### ゼロ値
goの変数は必ず初期化され、型それぞれの初期値が代入される
https://qiita.com/tenntenn/items/c55095585af64ca28ab5


### わかったこと
* 変数名などの名前を見れば可視性がわかる
* コンパイルの必要性があるが、ホットリロードの仕組みがある
* タプルによりエラー（の種類）、結果の正否を保持できるため、ガード節で前提となるエラーを弾きやすい
    * メソッド内部でオブジェクト内容を変更し、エラーを返却するという仕組みで処理自体は手続き的に行えそう
* ポインタによりメモリ節約の手段がある
* ポインタにより参照渡しであることが明言化されるため、渡した引数に副作用があることを明確化しやすい

## 参考文献
* [Webアプリ初心者がGo言語でサーバサイド（1. 簡単なHTTPサーバの実装）](https://qiita.com/wsuzume/items/75d5c0cd2dd5a1963b9e)
* [Go Web プログラミング](https://astaxie.gitbooks.io/build-web-application-with-golang/ja/)
* [Go Doc](https://godoc.org/?q=main)

* [【Go言語】パッケージのimportについて整理した](https://qiita.com/ogady/items/0cedd3599c4dc13e9a95)
