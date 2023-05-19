# gomock

基本的にはコマンドではなく`go:generate`でファイル内にコマンドを書いて生成していく
mockフォルダを作ってその中に`mock_+ファイル名`でモックを作っていくと良さそう
関数がどういった引数を与えられた状態で呼び出されてその場合どういった返り値を返したら良いのかを定義する

## モードについて
- source mode, reflect modeがあるが基本的にはreflect modeで良さそう
- ただ基本的には大きな違いというのはない
- reflect modeはエクスポートされていないインターフェースのモックは生成されない
- **一方例ではsource modeの方が良い 取り扱いが簡単ということもあってsource modeで良いかも？**

## テーブルテスト
テーブルテストをする場合各テスト内で新しいコントローラーを使って初期化は関数を記述することによって行う