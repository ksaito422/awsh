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
- pkg
  - あらゆるところから参照されるパッケージ
- tests
  - テスト用
- Makefile
  - 各種コマンド
- main.go
  - エントリポイント
