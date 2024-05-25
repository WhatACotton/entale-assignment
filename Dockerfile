FROM golang:1.22.3

# ワークディレクトリの設定
WORKDIR /app

# ソースコードのコピー
COPY . .

# 依存関係のインストール
RUN go mod download

# アプリケーションのビルド
RUN go build -o myapp

# ポートの設定
EXPOSE 8080

# コマンドの設定
CMD ["./myapp"]