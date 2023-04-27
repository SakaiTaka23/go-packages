# swagger

[オペレーションのアノテーション一覧](https://github.com/swaggo/swag/blob/master/README.md#api-operation)

アノテーションをパースしてドキュメント化はswagが行っていてGin用のライブラリはあくまでその表示をGinサーバーでできるようにしているだけ

サーバー上でドキュメントを表示する必要がない場合は必要ない

## コマンド
- `swag init` ドキュメント生成
- `swag fmt` コメントのフォーマット

swagのバージョンによってエラーが出るのでバージョンを下げるようにする
https://github.com/swaggo/swag/issues/1564
