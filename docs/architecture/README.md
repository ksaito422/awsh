# 本プロジェクトのディレクトリ構成について

```
awsh/
├ cmd/mani.go
├ docs/
├ internal/
│ 	└ endpoints/
├ pkg/*
├ tests/
└ Makefile

```

- cmd/main.go
  - アプリのエントリーポイント
- docs
  - ドキュメントの格納
- internal
  - 外部公開しないパッケージ
- internal/endpoints
  - ルーティングなどの起動時の処理を担当
  - aws-sdk を使ったメインの処理はほとんどここから呼ばれる想定
- pkg
  - あらゆるところから参照されるパッケージ
  - メインの処理がほとんど配下に置かれる
- tests
  - テスト用
- Makefile
  - 各種エイリアスコマンド
