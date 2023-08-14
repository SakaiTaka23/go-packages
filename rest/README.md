# Rest

- コードファーストでOpenAPIのドキュメントを生成することができる
- https://github.com/swaggest/rest
- 参考: https://dev.to/vearutop/tutorial-developing-a-restful-api-with-go-json-schema-validation-and-openapi-docs-2490



## 使い方

- 関数に対して直接タグを入れ込んでいくことで記述していく
- ハンドラーに関しては`usecase.Interactor`を使って記述する
  - 存在するエラーやタグの記述

- 今回はフレームワークごと変えたけどginをそのまま使って記述していくこともできそう
  - https://github.com/swaggest/rest/blob/master/_examples/gingonic/main.go

