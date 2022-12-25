# 本プロジェクトのディレクトリ構成について

```
awsh/
├ main.go
├ docs/
├ internal/
│ 	└ endpoints/
├ pkg/*
├ mock/
├ testutil/
└ Makefile

```

- main.go
  - CLIのエントリーポイント
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
- mock
  - テストで使うmock(gomock)
- testutil
  - テストで使う関数
- Makefile
  - 各種エイリアスコマンド
