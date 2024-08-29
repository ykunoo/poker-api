# ベースイメージとしてGoの公式イメージを使用
FROM golang:1.20-alpine

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールの依存関係ファイルをコピー
COPY go.mod go.sum ./

# 依存関係を取得
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o poker-api main.go

# デフォルトで実行されるコマンド
CMD ["./poker-api"]

# アプリケーションがリッスンするポートを指定
EXPOSE 8080
