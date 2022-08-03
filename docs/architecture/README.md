# 本プロジェクトのディレクトリ構成について

```
awsh/
├ docs/
├ internal/
├ pkg/
├ tests/
├ Makefile
└ main.go
```

- docs
  - ドキュメントの格納
- internal
  - 外部公開しないパッケージ
- internal/route
  - 操作対象の AWS リソースを選択して返す
- internal/controller
  - route で選択したリソースに対して、操作アクションを選択して実行する
  - aws-sdk を使ったメインの処理はほとんどここから呼ばれる想定
- pkg
  - あらゆるところから参照されるパッケージ
- tests
  - テスト用
- Makefile
  - 各種コマンド
- main.go
  - エントリポイント
