# Base

## ドキュメント周り
- [公式ドキュメント](https://gin-gonic.com/)
- 各種関数は[dev](https://pkg.go.dev/github.com/gin-gonic/gin@v1.9.0)の方をみる必要がありそう

## 立ち上げ
- `gin.Default`でロガーとリカバリーミドルウェアが入った状態で開始される
- `gin.New`の場合は何もなしで開始
ミドルウェアはリカバリーと[zap形式のロガー](https://github.com/gin-contrib/zap)は欲しいかも
