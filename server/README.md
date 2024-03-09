# GraphQLサーバーをGoで実装してみた

## チュートリアルを試してみる
```sh
# goプロジェクトの作成
$ go mod init example

# tools.goを作成し、gqlgenをインストールする
$ printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
$ go mod tidy

# gqlgenの初期化(サンプルコードが生成される)
$ go run github.com/99designs/gqlgen init
$ go mod tidy
```