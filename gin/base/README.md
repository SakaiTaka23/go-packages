# Base

## ドキュメント周り
- [公式ドキュメント](https://gin-gonic.com/)
- 各種関数は[dev](https://pkg.go.dev/github.com/gin-gonic/gin@v1.9.0)の方をみる必要がありそう

## 立ち上げ
- `gin.Default`でロガーとリカバリーミドルウェアが入った状態で開始される
- `gin.New`の場合は何もなしで開始
ミドルウェアはリカバリーと[zap形式のロガー](https://github.com/gin-contrib/zap)は欲しいかも

## ハンドラー
- `gin.HandlerFunc`を返してから行う方法もあるがハンドラーの場合は直接書いたので良さそう
- ミドルウェアを自作する場合は`gin.HandlerFunc`を使う
- ルート周りはハンドラーインスタンスの初期化時にやったので良さそう
