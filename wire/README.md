# Wire

## コード生成

```shell
wire wire/wire.go
```

## 実装の流れ

- 基本的には`wire.build`に依存する値を書いて行ってコード生成をするだけ
- 返り値は返して欲しいインスタンスの空の構造体にする
- その関数に対して引数が必要な場合はインスタンス関数に引数を渡すことによって解決する
  - これは型によって引数解決をしてそう？
  - 同じ型がある場合間違ったコードが生成される可能性があるのでstringなど基本的な型に依存している場合は専用の型エイリアスを作って上げる
- `wire.build`内に引数が過剰になっている場合や不足している場合はエラーを出してくれる
- ライブラリ作成中などでユーザーの状況によってインスタンス作成に必要な引数が変わる場合`wire.NewSet`でひとまとまりのセットを作っておくこともできる
  - 状況に応じて引数を変えてみる
