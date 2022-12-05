# go-api-e2e-test

このレポジトリは Qiitaに公開した記事 [runnを使ったWebバックエンドのe2eテスト](https://qiita.com/kourin1996/items/5c644210e12764d2caa3) のサンプルコードになります。

This is a sample project of e2e test for Go API, described in https://qiita.com/kourin1996/items/5c644210e12764d2caa3.  

Webフレームワークは [echo](https://github.com/labstack/echo)、e2eテストツールとして[runn](https://github.com/k1LoW/runn)を使用しています。

## How to start

```bash
$ go run main.go
```

```bash
$ go test ./...
```
