# Docker

https://pkg.go.dev/github.com/docker/docker/client

## インストール
- バージョンに注意すること
    - [Develop with Docker Engine API | Docker Documentation](https://docs.docker.com/engine/api/#versioned-api-and-sdk)
    - clientのバージョンがDocker本体より新しい場合は動かない
    - 必要に応じて`go get github.com/docker/docker/client`でバージョン指定する

## 例
https://docs.docker.com/engine/api/sdk/examples/

## ポート割り振り
- ポートを割り振る時にホスト側のポートを空文字にして送るとポートが自動的に割り振られるようになる
- 実際に開いたポートを確認するには関数を使ってIDからコンテナを確認してその中でポート情報を取得する
  - コンテナ一つの同じポートからホスト側から複数のポートを開くことができるため配列型になっていることに注意
