# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

# Claude Code 向けガイドライン（日本語訳）

このファイルは、Claude Code（claude.ai/code）がこのリポジトリ内のコードを扱う際のガイドラインを提供します。

## ビルドおよびテストコマンド

- **ビルド**: `go build -o taskai-cli main.go`
- **実行**: `./taskai-cli`
- **テスト**: `go test ./...`
- **単一パッケージのテスト**: `go test ./path/to/package`
- **単一関数のテスト**: `go test -run TestFunctionName ./path/to/package`
- **コードの整形**: `go fmt ./...`
- **リント**: `go vet ./...`

## コードスタイルガイドライン

- **整形**: `go fmt` によるGoの標準フォーマットに従う
- **インポート**: インポートはグループ分けする（標準ライブラリ → 外部 → 内部の順）
- **型**: 型には説明的な名前を使用し、`const` ブロックを使って強い型付け（列挙型）を推奨
- **命名**: 変数はキャメルケース（camelCase）、エクスポートされた型や関数はパスカルケース（PascalCase）を使用
- **エラー処理**: エラーは常にチェックし、適切に呼び出し元に返す
- **ドキュメント**: エクスポートされた関数には Go Doc 形式のコメントを付ける
- **ファイル構成**: 標準的なGoプロジェクトのレイアウト（`cmd/`, `internal/`, `pkg/`）に従う
- **テスト**: `*_test.go` ファイル内にテーブル駆動のテストを書く
- **依存関係**: 外部依存は最小限にとどめ、Go Modules で依存管理を行う
